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

type posMap map[pos]int

func neighbors() []pos {
	return []pos{pos{0, 1}, pos{1, 0}, pos{0, -1}, pos{-1, 0}}
}

func getRiskPath(m posMap, n int) int {
	risks := posMap{
		pos{0, 0}: 0,
	}
	next := []pos{pos{0, 0}}
	for len(next) > 0 {
		current := next[0]
		cr, _ := risks[current]
		next = next[1:]
		for _, n := range neighbors() {
			pn := pos{current.x + n.x, current.y + n.y}
			r, ok := m[pn]
			if ok {
				rr, already := risks[pn]
				if !already || rr > r+cr {
					risks[pn] = r + cr
					next = append(next, pn)
				}
			}
		}
	}
	risk, _ := risks[pos{n - 1, n - 1}]
	return risk
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := map[pos]int{}
	j := 0
	for scanner.Scan() {
		for i, ch := range []rune(scanner.Text()) {
			p := pos{i, j}
			v, _ := strconv.Atoi(string(ch))
			m[p] = v
		}
		j++
	}
	r1 := getRiskPath(m, j)
	fmt.Println("Part 1:", r1)
	for ty := 0; ty < 5; ty++ {
		for tx := 0; tx < 5; tx++ {
			cc := tx + ty
			for y := 0; y < j; y++ {
				for x := 0; x < j; x++ {
					orig := pos{x, y}
					target := pos{tx*j + x, ty*j + y}
					if orig == target {
						continue
					}
					value, _ := m[orig]
					for xx := 0; xx < cc; xx++ {
						if value == 9 {
							value = 1
						} else {
							value++
						}
					}
					m[target] = value
				}
			}
		}
	}
	r2 := getRiskPath(m, j*5)
	fmt.Println("Part 2:", r2)
}
