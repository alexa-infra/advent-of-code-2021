package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

type pos struct {
	x, y int
}

type posMap map[pos]int

func square3x3() []pos {
	return []pos{pos{-1, -1}, pos{0, -1}, pos{1, -1}, pos{-1, 0}, pos{0, 0}, pos{1, 0}, pos{-1, 1}, pos{0, 1}, pos{1, 1}}
}

func squareToInt(m posMap, p pos, defaultValue int) int {
	bin := make([]rune, 0, 9)
	for _, n := range square3x3() {
		pp := pos{p.x + n.x, p.y + n.y}
		v, ok := m[pp]
		if !ok {
			v = defaultValue
		}
		if v == 0 {
			bin = append(bin, '0')
		} else if v == 1 {
			bin = append(bin, '1')
		}
	}
	b := new(big.Int)
	b.SetString(string(bin), 2)
	return int(b.Int64())
}

func step(m posMap, h map[int]int, defaultValue int) posMap {
	nm := posMap{}
	for p, _ := range m {
		for _, n := range square3x3() {
			pp := pos{p.x + n.x, p.y + n.y}
			_, ok := nm[pp]
			if ok {
				continue
			}
			intVal := squareToInt(m, pp, defaultValue)
			vv, _ := h[intVal]
			nm[pp] = vv
		}
	}
	return nm
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	h := map[int]int{}
	for i, ch := range scanner.Text() {
		if ch == '.' {
			h[i] = 0
		} else if ch == '#' {
			h[i] = 1
		}
	}
	scanner.Scan()
	m := posMap{}
	j := 0
	for scanner.Scan() {
		for i, ch := range scanner.Text() {
			p := pos{i, j}
			if ch == '.' {
				m[p] = 0
			} else if ch == '#' {
				m[p] = 1
			}
		}
		j++
	}
	zero, _ := h[0]
	hasInvert := zero == 1
	for i := 0; i < 2; i++ {
		d := 0
		if hasInvert && i%2 != 0 {
			d = 1
		}
		m = step(m, h, d)
	}
	r1 := 0
	for _, v := range m {
		if v == 1 {
			r1++
		}
	}
	fmt.Println("Part 1:", r1)
	for i := 0; i < 48; i++ {
		d := 0
		if hasInvert && i%2 != 0 {
			d = 1
		}
		m = step(m, h, d)
	}
	r2 := 0
	for _, v := range m {
		if v == 1 {
			r2++
		}
	}
	fmt.Println("Part 2:", r2)
}
