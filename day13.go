package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type posMap map[pos]int

type foldLine struct {
	coord string
	val   int
}

func fold(m posMap, line foldLine) posMap {
	r := posMap{}
	if line.coord == "x" {
		for p := range m {
			if p.x > line.val {
				diff := p.x - line.val
				x := line.val - diff
				pp := pos{x, p.y}
				r[pp] = 1
			} else {
				r[p] = 1
			}
		}
	} else if line.coord == "y" {
		for p := range m {
			if p.y > line.val {
				diff := p.y - line.val
				y := line.val - diff
				pp := pos{p.x, y}
				r[pp] = 1
			} else {
				r[p] = 1
			}
		}
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := posMap{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 2 {
			break
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := pos{x, y}
		m[p] = 1
	}
	lines := []foldLine{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		defs := strings.Split(parts[2], "=")
		val, _ := strconv.Atoi(defs[1])
		line := foldLine{defs[0], val}
		lines = append(lines, line)
	}
	m1 := fold(m, lines[0])
	r1 := len(m1)
	fmt.Println("Part 1:", r1)
	for _, line := range lines {
		m = fold(m, line)
	}
	min, max := pos{0, 0}, pos{0, 0}
	first := true
	for p := range m {
		if first {
			min = p
			max = p
			first = false
		} else {
			if p.x > max.x {
				max.x = p.x
			}
			if p.y > max.y {
				max.y = p.y
			}
			if p.x < min.x {
				min.x = p.x
			}
			if p.y < min.y {
				min.y = p.y
			}
		}
	}
	fmt.Println("Part 2:")
	for j := min.y; j <= max.y; j++ {
		for i := min.x; i <= max.x; i++ {
			p := pos{i, j}
			_, ok := m[p]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
