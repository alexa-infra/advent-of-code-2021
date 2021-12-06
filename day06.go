package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func step(fishes []int) []int {
	new_fishes := 0
	for i, fish := range fishes {
		if fish > 0 {
			fishes[i] = fish - 1
		} else {
			fishes[i] = 6
			new_fishes++
		}
	}
	for i := 0; i < new_fishes; i++ {
		fishes = append(fishes, 8)
	}
	return fishes
}

func part1(inFishes []int, days int) int {
	fishes := make([]int, len(inFishes))
	copy(fishes, inFishes)
	for i := 0; i < days; i++ {
		fishes = step(fishes)
	}
	return len(fishes)
}

func part2(fishes []int, days int) int {
	groups := map[int]int{}
	for _, fish := range fishes {
		group, ok := groups[fish]
		if !ok {
			groups[fish] = 1
		} else {
			groups[fish] = group + 1
		}
	}
	for i := 0; i < days; i++ {
		new_groups := map[int]int{}
		for j := 0; j < 9; j++ {
			group, ok := groups[j]
			if !ok {
				group = 0
			}
			new_id := j - 1
			if new_id < 0 {
				new_id = 6
				new_groups[8] = group
			}
			new_group, ok := new_groups[new_id]
			if !ok {
				new_group = 0
			}
			new_groups[new_id] = new_group + group
		}
		groups = new_groups
	}
	n := 0
	for _, group := range groups {
		n += group
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")
	fishes := make([]int, 0, len(parts))
	for _, part := range parts {
		fish, _ := strconv.Atoi(part)
		fishes = append(fishes, fish)
	}
	r1 := part1(fishes, 80)
	fmt.Println("Part 1:", r1)
	r2 := part2(fishes, 256)
	fmt.Println("Part 2:", r2)
}
