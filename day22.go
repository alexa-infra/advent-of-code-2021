package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x, y, z int
}

type cube struct {
	xmin, xmax int
	ymin, ymax int
	zmin, zmax int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func intersect(a, b cube) (cube, bool) {
	if a.xmin <= b.xmin && a.xmax >= b.xmax && a.ymin <= b.ymin && a.ymax >= b.ymax && a.zmin <= b.zmin && a.zmax >= b.zmax {
		return cube{b.xmin, b.xmax, b.ymin, b.ymax, b.zmin, b.zmax}, true
	}

	if a.xmin >= b.xmin && a.xmax <= b.xmax && a.ymin >= b.ymin && a.ymax <= b.ymax && a.zmin >= b.zmin && a.zmax <= b.zmax {
		return cube{a.xmin, a.xmax, a.ymin, a.ymax, a.zmin, a.zmax}, true
	}

	if a.xmin > b.xmax || a.ymin > b.ymax || a.zmin > b.zmax || a.xmax < b.xmin || a.ymax < b.ymin || a.zmax < b.zmin {
		return cube{}, false
	}

	return cube{max(a.xmin, b.xmin), min(a.xmax, b.xmax), max(a.ymin, b.ymin), min(a.ymax, b.ymax), max(a.zmin, b.zmin), min(a.zmax, b.zmax)}, true
}

func (c cube) size() int {
	return (c.xmax + 1 - c.xmin) * (c.ymax + 1 - c.ymin) * (c.zmax + 1 - c.zmin)
}

func atoi(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("wrong num format")
	}
	return val
}

type step struct {
	on   bool
	area cube
}

func main() {
	re := regexp.MustCompile(`(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	steps := []step{}
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		if match == nil {
			log.Fatalf("wrong format")
		}
		on := false
		if match[1] == "on" {
			on = true
		}
		xmin, xmax := atoi(match[2]), atoi(match[3])
		ymin, ymax := atoi(match[4]), atoi(match[5])
		zmin, zmax := atoi(match[6]), atoi(match[7])
		c := cube{xmin, xmax, ymin, ymax, zmin, zmax}
		steps = append(steps, step{on, c})
	}
	m := map[coord]int{}
	initArea := cube{-50, 50, -50, 50, -50, 50}
	for _, s := range steps {
		on := s.on
		area := s.area
		for z := max(initArea.zmin, area.zmin); z <= min(initArea.zmax, area.zmax); z++ {
			for y := max(initArea.ymin, area.ymin); y <= min(initArea.ymax, area.ymax); y++ {
				for x := max(initArea.xmin, area.xmin); x <= min(initArea.xmax, area.xmax); x++ {
					p := coord{x, y, z}
					_, ok := m[p]
					if !ok && on {
						m[p] = 1
					} else if ok && !on {
						delete(m, p)
					}
				}
			}
		}
	}
	r1 := len(m)
	fmt.Println("Part 1:", r1)

	new_steps := []step{}
	for _, s := range steps {
		extra := []step{}
		on, area := s.on, s.area
		if on {
			extra = append(extra, step{on, area})
		}
		for _, prev := range new_steps {
			new_area, ok := intersect(area, prev.area)
			if !ok {
				continue
			}
			extra = append(extra, step{!prev.on, new_area})
		}
		new_steps = append(new_steps, extra...)
	}
	r2 := 0
	for _, s := range new_steps {
		if s.on {
			r2 += s.area.size()
		} else {
			r2 -= s.area.size()
		}
	}
	fmt.Println("Part 2:", r2)
}
