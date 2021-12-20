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

func squareToInt(m posMap, p pos, inv bool) int {
	bin := make([]rune, 0, 9)
	for _, n := range square3x3() {
		pp := pos{p.x + n.x, p.y + n.y}
		v, ok := m[pp]
		if !ok {
			v = 0
		}
		if !inv {
			if v == 0 {
				bin = append(bin, '0')
			} else if v == 1 {
				bin = append(bin, '1')
			}
		} else {
			if v == 0 {
				bin = append(bin, '1')
			} else if v == 1 {
				bin = append(bin, '0')
			}
		}
	}
	b := new(big.Int)
	b.SetString(string(bin), 2)
	return int(b.Int64())
}

func step(m posMap, h map[int]int, inv bool) posMap {
	nm := posMap{}
	for p, v := range m {
		if v != 1 {
			continue
		}
		for _, n := range square3x3() {
			pp := pos{p.x + n.x, p.y + n.y}
			_, ok := nm[pp]
			if ok {
				continue
			}
			intVal := squareToInt(m, pp, inv)
			vv, _ := h[intVal]
			nm[pp] = vv
		}
	}
	return nm
}

func invertMap(m posMap) posMap {
	n := posMap{}
	for p, v := range m {
		if v == 1 {
			n[p] = 0
		} else if v == 0 {
			n[p] = 1
		}
	}
	return n
}

func invertH(h map[int]int) map[int]int {
	hh := map[int]int{}
	for k, v := range h {
		if v == 1 {
			hh[k] = 0
		} else if v == 0 {
			hh[k] = 1
		}
	}
	return hh
}

func pprint(m posMap) {
	min, max := pos{}, pos{}
	first := true
	for p, _ := range m {
		if first {
			min = p
			max = p
			first = false
		} else {
			if p.x < min.x {
				min.x = p.x
			}
			if p.y < min.y {
				min.y = p.y
			}
			if p.x > max.x {
				max.x = p.x
			}
			if p.y > max.y {
				max.y = p.y
			}
		}
	}
	for j := min.y; j <= max.y; j++ {
		for i := min.x; i <= max.x; i++ {
			p := pos{i, j}
			v, ok := m[p]
			if !ok {
				v = 0
			}
			if v == 1 {
				fmt.Printf("#")
			} else if v == 0 {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println()
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
	if hasInvert {
		for i := 0; i < 2; i++ {
			m = step(m, h, i%2 != 0)
			m = invertMap(m)
			h = invertH(h)
		}
	} else {
		for i := 0; i < 2; i++ {
			m = step(m, h, false)
		}
	}
	r1 := 0
	for _, v := range m {
		if v == 1 {
			r1++
		}
	}
	fmt.Println("Part 1:", r1)
	if hasInvert {
		for i := 0; i < 48; i++ {
			m = step(m, h, i%2 != 0)
			m = invertMap(m)
			h = invertH(h)
		}
	} else {
		for i := 0; i < 48; i++ {
			m = step(m, h, false)
		}
	}
	r2 := 0
	for _, v := range m {
		if v == 1 {
			r2++
		}
	}
	fmt.Println("Part 2:", r2)
}
