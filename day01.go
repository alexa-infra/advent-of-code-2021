package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	depths := []int{}
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}

	prevDepth := 0
	first := true
	r1 := 0
	for _, depth := range depths {
		if !first && depth > prevDepth {
			r1++
		}
		if first {
			first = false
		}
		prevDepth = depth
	}
	fmt.Println("Part 1:", r1)

	wLen := 3
	prevWinSum := 0
	first = true
	r2 := 0
	for i := 0; i < len(depths)-wLen+1; i++ {
		window := depths[i : i+wLen]
		sum := 0
		for _, depth := range window {
			sum += depth
		}
		if !first && sum > prevWinSum {
			r2++
		}
		if first {
			first = false
		}
		prevWinSum = sum
	}
	fmt.Println("Part 2:", r2)
}
