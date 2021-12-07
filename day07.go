package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func fuelToMove(positions []int, center int) int {
	fuel := 0
	for _, pos := range positions {
		fuel += abs(center - pos)
	}
	return fuel
}

func fuelToMove2(positions []int, center int) int {
	fuel := 0
	for _, pos := range positions {
		for i := 0; i < abs(center-pos); i++ {
			fuel += i + 1
		}
	}
	return fuel
}

type fuelFunc func(positions []int, center int) int

func minFuelToMove(positions []int, f fuelFunc) int {
	s := 0
	for _, pos := range positions {
		s += pos
	}
	s /= len(positions)
	left := f(positions, s-1)
	sf := f(positions, s)
	right := f(positions, s+1)
	if left < sf {
		for left < sf {
			s--
			left = f(positions, s-1)
			sf = f(positions, s)
		}
	} else if right < sf {
		for right < sf {
			s++
			sf = f(positions, s)
			right = f(positions, s+1)
		}
	}
	return sf
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")
	positions := []int{}
	for _, part := range parts {
		pos, _ := strconv.Atoi(part)
		positions = append(positions, pos)
	}
	r1 := minFuelToMove(positions, fuelToMove)
	fmt.Println("Part 1:", r1)
	r2 := minFuelToMove(positions, fuelToMove2)
	fmt.Println("Part 2:", r2)
}
