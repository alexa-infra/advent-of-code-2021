package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func step(pairs map[string]int, convert map[string]string) map[string]int {
	result := map[string]int{}
	for k, v := range pairs {
		r, ok := convert[k]
		if !ok {
			result[k] = v
		} else {
			elements := []rune(k)
			el1, el2 := elements[0], elements[1]
			results := []rune(r)
			rel := results[0]
			p1 := string([]rune{el1, rel})
			p2 := string([]rune{rel, el2})
			vv, kk := result[p1]
			if !kk {
				vv = 0
			}
			result[p1] = vv + v
			vv, kk = result[p2]
			if !kk {
				vv = 0
			}
			result[p2] = vv + v
		}
	}
	return result
}

func count(pairs map[string]int) int {
	cc := map[rune]int{}
	for k, v := range pairs {
		elements := []rune(k)
		el := elements[0]
		vv, ok := cc[el]
		if !ok {
			vv = 0
		}
		cc[el] = vv + v
	}
	min, max := 0, 0
	first := true
	for _, v := range cc {
		if first {
			min = v
			max = v
			first = false
		} else {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}
	return max - min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	n := len(line)
	pairs := map[string]int{}
	for i := 0; i < n-1; i++ {
		pair := line[i : i+2]
		v, ok := pairs[pair]
		if !ok {
			v = 0
		}
		pairs[pair] = v + 1
	}
	tail := line[n-1:]
	pairs[tail] = 1
	scanner.Scan()
	convert := map[string]string{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		pair, result := parts[0], parts[2]
		convert[pair] = result
	}
	m1 := pairs
	for i := 0; i < 10; i++ {
		m1 = step(m1, convert)
	}
	r1 := count(m1)
	fmt.Println("Part 1:", r1)
	m2 := pairs
	for i := 0; i < 40; i++ {
		m2 = step(m2, convert)
	}
	r2 := count(m2)
	fmt.Println("Part 2:", r2)
}
