package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y int
}

type board map[pos]rune

var energy map[rune]int = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

var finalRoomX map[rune]int = map[rune]int{'A': 2, 'B': 4, 'C': 6, 'D': 8}

var hallPositions []pos = []pos{
	pos{0, 0},
	pos{1, 0},
	pos{3, 0},
	pos{5, 0},
	pos{7, 0},
	pos{9, 0},
	pos{10, 0},
}

func isFinal(state board, depth int) bool {
	// check if all pockets are filled with correct balls
	for _, x := range finalRoomX {
		for d := 1; d <= depth; d++ {
			p := pos{x, d}
			if !isFinalPos(p, state, depth) {
				return false
			}
		}
	}
	return true
}

func isFinalPos(p pos, state board, depth int) bool {
	// check if position is filled, position is in the pocket, the pocket is correct
	// and positions below are also correctly filled
	ch, ok := state[p]
	if !ok {
		return false
	}
	finalX, _ := finalRoomX[ch]
	if p.x != finalX {
		return false
	}
	for d := depth; d >= p.y; d-- {
		pp := pos{finalX, d}
		v, ok := state[pp]
		if !ok || v != ch {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type path struct {
	start, dest pos
}

func (p path) distance() int {
	// mahnhatten distance
	return abs(p.dest.x-p.start.x) + abs(p.dest.y-p.start.y)
}

func getEnergy(ch rune, p path) int {
	en, _ := energy[ch]
	return en * p.distance()
}

func getPaths(state board, depth int) []path {
	// returns all possible moves/paths of the current state
	ret := []path{}
	for p, _ := range state {
		dest := getDestinations(p, state, depth)
		for _, d := range dest {
			ret = append(ret, path{p, d})
		}
	}
	return ret
}

func getDestinations(p pos, state board, depth int) []pos {
	// returns all possible moves/paths of the current position/ball
	ret := []pos{}
	if p.y != 0 {
		// the ball is inside the pocket
		if isFinalPos(p, state, depth) {
			// the ball is inside its target pocket, no need to move it
			return ret
		}
		for d := p.y - 1; d > 0; d-- {
			// check if above places in the pocket are empty
			// if not-empty is found, then we can't move the ball
			_, ok := state[pos{p.x, d}]
			if ok {
				return ret
			}
		}
		for i := 0; i < len(hallPositions); i++ {
			// collect empty spaces in the hall to the right from the pocket
			h := hallPositions[i]
			if h.x > p.x {
				_, ok := state[h]
				if ok {
					break
				}
				ret = append(ret, h)
			}
		}
		for i := len(hallPositions) - 1; i >= 0; i-- {
			// collect empty spaces in the hall to the left from the pocket
			h := hallPositions[i]
			if h.x < p.x {
				_, ok := state[h]
				if ok {
					break
				}
				ret = append(ret, h)
			}
		}
	} else {
		// the ball is in the hall, so we can move it only inside its target pocket
		cc, _ := state[p]
		finalX, _ := finalRoomX[cc]
		obstacle := false
		if finalX > p.x {
			// we move from right to left in the hall and check if there are obstacles
			for i := 0; i < len(hallPositions); i++ {
				h := hallPositions[i]
				if h.x > p.x && h.x < finalX {
					_, ok := state[h]
					if ok {
						obstacle = true
						break
					}
				}
			}
		} else {
			// we move from left to right in the hall and check if there are obstacles
			for i := len(hallPositions) - 1; i >= 0; i-- {
				h := hallPositions[i]
				if h.x < p.x && h.x > finalX {
					_, ok := state[h]
					if ok {
						obstacle = true
						break
					}
				}
			}
		}
		if !obstacle {
			// there are no obstacles in the hall
			// so we check if we can move to the pocket from top to bottom
			for d := depth; d > 0; d-- {
				room := pos{finalX, d}
				vv, kk := state[room]
				if !kk {
					// the first empty space, path is clear, move!
					ret = append(ret, room)
					break
				} else if vv != cc {
					// we can't move if space is taken by non-target ball
					break
				}
			}
		}
	}
	return ret
}

type step struct {
	n     int
	g     int
	state board
}

func toString(state board, depth int) string {
	s := ""
	for j := 0; j <= depth; j++ {
		for i := 0; i <= 10; i++ {
			ch, ok := state[pos{i, j}]
			if !ok {
				s += "."
			} else {
				s += string(ch)
			}
		}
		s += "\n"
	}
	return s
}

func getDepth(m board) int {
	depth := 0
	for p, _ := range m {
		if p.y > depth {
			depth = p.y
		}
	}
	return depth
}

func solve(m board) int {
	depth := getDepth(m)
	steps := []step{
		step{0, 0, m},
	}
	r1 := 0
	first := true
	cache := map[string]int{}
	for len(steps) > 0 {
		current := steps[0]
		steps = steps[1:]
		paths := getPaths(current.state, depth)
		if !first && current.n > r1 {
			// if we've already found one solution, then we can reject
			// half-way paths by their current energy
			continue
		}
		cacheKey := toString(current.state, depth)
		ccc, kkk := cache[cacheKey]
		if kkk {
			// if we've already seen this state before and it was lower
			// by the energy, then there are no reasons to continue with
			// greater energy
			if ccc <= current.n {
				continue
			}
		}
		cache[cacheKey] = current.n
		if len(paths) == 0 {
			// no possible moves = we should check if we filled all pockets
			// so it's a solution of not
			if isFinal(current.state, depth) {
				if first {
					r1 = current.n
					first = false
				} else if current.n < r1 {
					r1 = current.n
				}
			}
		} else {
			// for each possible move we extend next steps
			for _, p := range paths {
				new_state := board{}
				ch, _ := current.state[p.start]
				for k, v := range current.state {
					if k != p.start {
						new_state[k] = v
					}
				}
				new_state[p.dest] = ch
				new_energy := current.n + getEnergy(ch, p)
				steps = append(steps, step{new_energy, current.g + 1, new_state})
			}
		}
	}
	return r1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	j := 0
	m := board{}
	for scanner.Scan() {
		line := scanner.Text()
		for i, ch := range line {
			if ch == 'A' || ch == 'B' || ch == 'C' || ch == 'D' {
				p := pos{i - 1, j - 1}
				m[p] = ch
			}
		}
		j++
	}
	r1 := solve(m)
	fmt.Println("Path 1:", r1)
}
