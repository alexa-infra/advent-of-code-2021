package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type pos struct {
	x, y int
}

func diffs() []pos {
	return []pos{pos{0, 1}, pos{0, -1}, pos{1, 0}, pos{-1, 0}}
}

type posMap map[pos]int

func (p *posMap) neighbors(d pos) []pos {
	ret := []pos{}
	for _, diff := range diffs() {
		dd := pos{d.x + diff.x, d.y + diff.y}
		_, ok := (*p)[dd]
		if ok {
			ret = append(ret, dd)
		}
	}
	return ret
}

func (p *posMap) flowToLowPoint(d pos, basins *posMap) int {
	t := d
	v, _ := (*p)[t]
	if v == 9 {
		return -1
	}
	path := []pos{}
	for {
		basinId, ok := (*basins)[t]
		if ok {
			for _, pp := range path {
				(*basins)[pp] = basinId
			}
			return basinId
		} else {
			path = append(path, t)
		}
		v, _ = (*p)[t]
		neighbors := p.neighbors(t)
		minNeighbor := 9
		for _, n := range neighbors {
			nn, _ := (*p)[n]
			if nn < minNeighbor {
				minNeighbor = nn
				t = n
			}
		}
	}
	return -1
}

func minArr(arr []int) int {
	m := 0
	first := true
	for _, v := range arr {
		if first || v < m {
			m = v
			first = false
		}
	}
	return m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := posMap{}
	j := 0
	for scanner.Scan() {
		for i, ch := range scanner.Text() {
			p := pos{i, j}
			v, _ := strconv.Atoi(string(ch))
			m[p] = v
		}
		j++
	}
	r1 := 0
	basins := posMap{}
	basinId := 0
	for p, v := range m {
		neighbors := m.neighbors(p)
		minNeighbor := 9
		for _, n := range neighbors {
			nn, _ := m[n]
			if nn < minNeighbor {
				minNeighbor = nn
			}
		}
		if v < minNeighbor {
			r1 += v + 1
			basins[p] = basinId
			basinId++
		}
	}
	fmt.Println("Part 1:", r1)
	for p, _ := range m {
		m.flowToLowPoint(p, &basins)
	}
	basinCounts := map[int]int{}
	for _, i := range basins {
		cc, ok := basinCounts[i]
		if !ok {
			cc = 0
		}
		basinCounts[i] = cc + 1
	}
	basinCountsArr := make([]int, len(basinCounts))
	for i, v := range basinCounts {
		basinCountsArr[i] = v
	}
	arrSlice := basinCountsArr[:]
	sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))
	r2 := basinCountsArr[0] * basinCountsArr[1] * basinCountsArr[2]
	fmt.Println("Part 2:", r2)
}
