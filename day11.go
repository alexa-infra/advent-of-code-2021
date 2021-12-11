package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x, y int
}

func neighbors() []pos {
	return []pos{pos{0, -1}, pos{1, -1}, pos{1, 0}, pos{1, 1}, pos{0, 1}, pos{-1, 1}, pos{-1, 0}, pos{-1, -1}}
}

func step(m map[pos]int) int {
	flashing := []pos{}
	flashed := map[pos]int{}
	for p, v := range m {
		m[p] = v + 1
		if v+1 > 9 {
			flashing = append(flashing, p)
		}
	}
	for len(flashing) > 0 {
		p := flashing[0]
		flashing = flashing[1:]
		_, ok := flashed[p]
		if ok {
			continue
		}
		m[p] = 0
		flashed[p] = 0
		for _, n := range neighbors() {
			d := pos{p.x + n.x, p.y + n.y}
			v, ok := m[d]
			if !ok {
				continue
			}
			_, ok = flashed[d]
			if ok {
				continue
			}
			m[d] = v + 1
			if v+1 > 9 {
				flashing = append(flashing, d)
			}
		}
	}
	return len(flashed)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := map[pos]int{}
	j := 0
	for scanner.Scan() {
		for i, ch := range []rune(scanner.Text()) {
			v, _ := strconv.Atoi(string(ch))
			p := pos{i, j}
			m[p] = v
		}
		j++
	}
	n := len(m)
	r1 := 0
	r2 := 0
	i := 0
	for {
		nn := step(m)
		if i < 100 {
			r1 += nn
		}
		i++
		if nn == n {
			r2 = i
			break
		}
	}
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)
}
