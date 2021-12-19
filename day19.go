package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	x, y, z int
}

func diff(a, b vec) vec {
	return vec{a.x - b.x, a.y - b.y, a.z - b.z}
}

func add(a, b vec) vec {
	return vec{a.x + b.x, a.y + b.y, a.z + b.z}
}

type matrix struct {
	a0, a1, a2 int
	a3, a4, a5 int
	a6, a7, a8 int
}

func conv(m matrix, v vec) vec {
	r0 := m.a0*v.x + m.a1*v.y + m.a2*v.z
	r1 := m.a3*v.x + m.a4*v.y + m.a5*v.z
	r2 := m.a6*v.x + m.a7*v.y + m.a8*v.z
	return vec{r0, r1, r2}
}

func convArr(m matrix, v []vec) []vec {
	r := []vec{}
	for _, a := range v {
		b := conv(m, a)
		r = append(r, b)
	}
	return r
}

func moveArr(v []vec, d vec) []vec {
	r := []vec{}
	for _, x := range v {
		y := add(x, d)
		r = append(r, y)
	}
	return r
}

func cmp(a, b []vec) []vec {
	r := []vec{}
	for _, x := range a {
		for _, y := range b {
			if x == y {
				r = append(r, x)
			}
		}
	}
	return r
}

func join(a, b []vec) []vec {
	r := map[vec]int{}
	for _, x := range a {
		r[x] = 1
	}
	for _, y := range b {
		r[y] = 1
	}
	res := []vec{}
	for k, _ := range r {
		res = append(res, k)
	}
	return res
}

func mult(a, b matrix) matrix {
	r0 := a.a0*b.a0 + a.a1*b.a3 + a.a2*b.a6
	r1 := a.a0*b.a1 + a.a1*b.a4 + a.a2*b.a7
	r2 := a.a0*b.a2 + a.a1*b.a5 + a.a2*b.a8
	r3 := a.a3*b.a0 + a.a4*b.a3 + a.a5*b.a6
	r4 := a.a3*b.a1 + a.a4*b.a4 + a.a5*b.a7
	r5 := a.a3*b.a2 + a.a4*b.a5 + a.a5*b.a8
	r6 := a.a6*b.a0 + a.a7*b.a3 + a.a8*b.a6
	r7 := a.a6*b.a1 + a.a7*b.a4 + a.a8*b.a7
	r8 := a.a6*b.a2 + a.a7*b.a5 + a.a8*b.a8
	return matrix{
		r0, r1, r2,
		r3, r4, r5,
		r6, r7, r8,
	}
}

func permA() []matrix {
	return []matrix{
		matrix{1, 0, 0,
			0, 1, 0,
			0, 0, 1},
		matrix{0, 1, 0,
			0, 0, 1,
			1, 0, 0},
		matrix{0, 0, 1,
			1, 0, 0,
			0, 1, 0},
	}
}

func permB() []matrix {
	return []matrix{
		matrix{1, 0, 0,
			0, 1, 0,
			0, 0, 1},
		matrix{-1, 0, 0,
			0, -1, 0,
			0, 0, 1},
		matrix{-1, 0, 0,
			0, 1, 0,
			0, 0, -1},
		matrix{1, 0, 0,
			0, -1, 0,
			0, 0, -1},
	}
}

func permC() []matrix {
	return []matrix{
		matrix{1, 0, 0,
			0, 1, 0,
			0, 0, 1},
		matrix{0, 0, -1,
			0, -1, 0,
			-1, 0, 0},
	}
}

func overlap(sc []vec, res []vec) (bool, vec, []vec) {
	for _, c := range permC() {
		for _, b := range permB() {
			for _, a := range permA() {
				m := mult(c, mult(b, a))
				sc1 := convArr(m, sc)
				for _, f1 := range sc1 {
					for _, f2 := range res {
						dd := diff(f2, f1)
						sc2 := moveArr(sc1, dd)

						cc := cmp(res, sc2)
						if len(cc) >= 12 {
							return true, dd, sc2
						}
					}
				}
			}
		}
	}
	return false, vec{}, nil
}

type pair struct {
	i, j int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(a, b vec) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	beacons := []vec{}
	scanners := [][]vec{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "scanner") {
			beacons = []vec{}
		} else if line == "" {
			scanners = append(scanners, beacons)
		} else {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			v := vec{x, y, z}
			beacons = append(beacons, v)
		}
	}
	scanners = append(scanners, beacons)

	known := map[int]int{}
	known[0] = 0
	checked := map[pair]int{}

	pos := []vec{vec{0, 0, 0}}
	for len(known) != len(scanners) {
		for i, dest := range scanners {
			_, ok := known[i]
			if ok {
				continue
			}
			found := false
			for j := range known {
				pp := pair{i, j}
				_, ok := checked[pp]
				if ok {
					continue
				}
				src := scanners[j]
				ok, d, nv := overlap(dest, src)
				checked[pp] = 1
				if ok {
					pos = append(pos, d)
					scanners[i] = nv
					known[i] = 1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	r := []vec{}
	for _, k := range scanners {
		r = join(r, k)
	}
	fmt.Println("Part 1:", len(r))
	m := 0
	for i, x := range pos {
		for j, y := range pos {
			if i == j {
				continue
			}
			d := distance(x, y)
			if d > m {
				m = d
			}
		}
	}
	fmt.Println("Part 2:", m)
}
