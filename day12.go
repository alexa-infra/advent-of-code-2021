package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	connections := map[string][]string{}
	cavesize := map[string]bool{}
	for scanner.Scan() {
		connection := scanner.Text()
		parts := strings.Split(connection, "-")
		cave1, cave2 := parts[0], parts[1]
		cc, ok := connections[cave1]
		if !ok {
			cc = []string{}
		}
		connections[cave1] = append(cc, cave2)
		cc, ok = connections[cave2]
		if !ok {
			cc = []string{}
		}
		connections[cave2] = append(cc, cave1)
		_, ok = cavesize[cave1]
		if !ok {
			isBig := strings.ToUpper(cave1) == cave1
			cavesize[cave1] = isBig
		}
		_, ok = cavesize[cave2]
		if !ok {
			isBig := strings.ToUpper(cave2) == cave2
			cavesize[cave2] = isBig
		}
	}
	paths := [][]string{
		[]string{"start"},
	}
	r1 := 0
	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		current := path[len(path)-1]
		cc, _ := connections[current]
		for _, c := range cc {
			isBig, _ := cavesize[c]
			if isBig {
				newpath := []string{}
				newpath = append(newpath, path...)
				newpath = append(newpath, c)
				paths = append(paths, newpath)
			} else if c == "start" {
				continue
			} else if c == "end" {
				r1++
			} else {
				found := false
				for _, p := range path {
					if p == c {
						found = true
						break
					}
				}
				if !found {
					newpath := []string{}
					newpath = append(newpath, path...)
					newpath = append(newpath, c)
					paths = append(paths, newpath)
				}
			}
		}
	}
	fmt.Println("Part 1:", r1)
	paths = [][]string{
		[]string{"start"},
	}
	r2 := 0
	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		current := path[len(path)-1]
		cc, _ := connections[current]
		for _, c := range cc {
			isBig, _ := cavesize[c]
			if isBig {
				newpath := []string{}
				newpath = append(newpath, path...)
				newpath = append(newpath, c)
				paths = append(paths, newpath)
			} else if c == "start" {
				continue
			} else if c == "end" {
				r2++
			} else {
				dups := map[string]int{}
				for _, p := range path {
					isBig, _ = cavesize[p]
					if !isBig {
						v, ok := dups[p]
						if !ok {
							v = 0
						}
						dups[p] = v + 1
					}
				}
				hasDup := false
				for _, v := range dups {
					if v > 1 {
						hasDup = true
					}
				}
				countC, hasC := dups[c]
				if !hasDup || (hasDup && !hasC) || (!hasDup && countC == 1) {
					newpath := []string{}
					newpath = append(newpath, path...)
					newpath = append(newpath, c)
					paths = append(paths, newpath)
				}
			}
		}
	}
	fmt.Println("Part 2:", r2)
}
