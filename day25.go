package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y int
}

type posMap map[pos]rune

func getDest(m posMap, p pos) pos {
	ch, _ := m[p]
	if ch == '.' {
		return p
	}
	if ch == 'v' {
		pp := pos{p.x, p.y + 1}
		_, ok := m[pp]
		if !ok {
			pp = pos{p.x, 0}
		}
		return pp
	}
	if ch == '>' {
		pp := pos{p.x + 1, p.y}
		_, ok := m[pp]
		if !ok {
			pp = pos{0, p.y}
		}
		return pp
	}
	return pos{0, 0}
}

var directions []rune = []rune{'>', 'v'}

func step(m posMap) int {
	n := 0
	for _, dir := range directions {
		canMove := map[pos]pos{}
		for p, ch := range m {
			if ch == dir {
				pp := getDest(m, p)
				if pp == p {
					continue
				}
				val, ok := m[pp]
				if ok && val == '.' {
					canMove[p] = pp
				}
			}
		}
		for p, pp := range canMove {
			m[pp] = dir
			m[p] = '.'
		}
		n += len(canMove)
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := posMap{}
	j := 0
	for scanner.Scan() {
		for i, ch := range scanner.Text() {
			p := pos{i, j}
			m[p] = ch
		}
		j++
	}
	r1 := 0
	for {
		r1++
		moved := step(m)
		if moved == 0 {
			break
		}
	}
	fmt.Println("Part 1:", r1)
}
