package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

type pointLine struct {
	a, b point
}

func main() {
	reLine := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	lines := []pointLine{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := reLine.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(parts[1])
		y1, _ := strconv.Atoi(parts[2])
		x2, _ := strconv.Atoi(parts[3])
		y2, _ := strconv.Atoi(parts[4])
		a := point{x1, y1}
		b := point{x2, y2}
		lines = append(lines, pointLine{a, b})
	}
	data := map[point]int{}
	setData := func(p point) {
		pp, ok := data[p]
		if !ok {
			data[p] = 1
		} else {
			data[p] = pp + 1
		}
	}
	drawLine := func(a, b point) {
		dx := 0
		if a.x < b.x {
			dx = 1
		} else if a.x > b.x {
			dx = -1
		}
		dy := 0
		if a.y < b.y {
			dy = 1
		} else if a.y > b.y {
			dy = -1
		}
		p := a
		for p != b {
			setData(p)
			p.x += dx
			p.y += dy
		}
		setData(p)
	}
	r1 := 0
	for _, line := range lines {
		if line.a.x == line.b.x || line.a.y == line.b.y {
			drawLine(line.a, line.b)
		}
	}
	for _, v := range data {
		if v > 1 {
			r1++
		}
	}
	fmt.Println("Part 1:", r1)
	data = map[point]int{}
	r2 := 0
	for _, line := range lines {
		drawLine(line.a, line.b)
	}
	for _, v := range data {
		if v > 1 {
			r2++
		}
	}
	fmt.Println("Part 2:", r2)
}
