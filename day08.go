package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
0: abcefg
1: cf
2: acdeg
3: acdfg
4: bcdf
5: abdfg
6: abdefg
7: acf
8: abcdefg
9: abcdfg
*/

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func mappingUnique() map[string]int {
	return map[string]int{
		"cf":      1,
		"acf":     7,
		"bcdf":    4,
		"abcdefg": 8,
	}
}

func mappingNotUnique() map[string]int {
	return map[string]int{
		"acdeg":  2,
		"acdfg":  3,
		"abdfg":  5,
		"abcefg": 0,
		"abdefg": 6,
		"abcdfg": 9,
	}
}

func replace(a string, remap map[rune]rune) string {
	b := []rune(a)
	r := []rune{}
	for _, ch := range b {
		x, _ := remap[ch]
		r = append(r, x)
	}
	return sortString(string(r))
}

func unique(arr []string) []string {
	r := map[string]int{}
	for _, s := range arr {
		r[s] = 1
	}
	ret := []string{}
	for s, _ := range r {
		ret = append(ret, s)
	}
	return ret
}

func diff(a string, p ...string) string {
	r := []rune{}
	for _, x := range []rune(a) {
		found := false
		for _, b := range p {
			for _, y := range []rune(b) {
				if x == y {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			r = append(r, x)
		}
	}
	return string(r)
}

type pair struct {
	x, y []rune
}

func newPair(x, y string) pair {
	return pair{[]rune(x), []rune(y)}
}

func getVariants(p0, p1, p2, p3 pair) []string {
	mm := map[rune]rune{
		'a': ' ',
		'b': ' ',
		'c': ' ',
		'd': ' ',
		'e': ' ',
		'f': ' ',
		'g': ' ',
	}
	mm[p0.x[0]] = p0.y[0]
	r := []string{}
	for i1 := 0; i1 < 2; i1++ {
		if i1 == 0 {
			mm[p1.x[0]] = p1.y[0]
			mm[p1.x[1]] = p1.y[1]
		} else if i1 == 1 {
			mm[p1.x[0]] = p1.y[1]
			mm[p1.x[1]] = p1.y[0]
		}
		for i2 := 0; i2 < 2; i2++ {
			if i2 == 0 {
				mm[p2.x[0]] = p2.y[0]
				mm[p2.x[1]] = p2.y[1]
			} else if i2 == 1 {
				mm[p2.x[0]] = p2.y[1]
				mm[p2.x[1]] = p2.y[0]
			}
			for i3 := 0; i3 < 2; i3++ {
				if i3 == 0 {
					mm[p3.x[0]] = p3.y[0]
					mm[p3.x[1]] = p3.y[1]
				} else if i3 == 1 {
					mm[p3.x[0]] = p3.y[1]
					mm[p3.x[1]] = p3.y[0]
				}
				runes := []rune{
					mm['a'],
					mm['b'],
					mm['c'],
					mm['d'],
					mm['e'],
					mm['f'],
					mm['g'],
				}
				r = append(r, string(runes))
			}
		}
	}
	return r
}

func decode(left, right []string) int {
	all := append(append([]string{}, left...), right...)
	byLen := map[int][]string{}
	for _, s := range all {
		n := len(s)
		arr, ok := byLen[n]
		if !ok {
			byLen[n] = []string{s}
		} else {
			byLen[n] = unique(append(arr, s))
		}
	}
	b1, ok := byLen[2]
	if !ok {
		log.Fatalf("missing 1")
	}
	v1 := b1[0]
	w1 := "cf"
	p1 := newPair(v1, w1)
	b7, ok := byLen[3]
	if !ok {
		log.Fatalf("missing 7")
	}
	v7 := b7[0]
	w7 := "acf"
	p0 := newPair(diff(v7, v1), diff(w7, w1))
	b4, ok := byLen[4]
	if !ok {
		log.Fatalf("missing 4")
	}
	v4 := b4[0]
	w4 := "bcdf"
	p2 := newPair(diff(v4, v1), diff(w4, w1))
	b8, ok := byLen[7]
	if !ok {
		log.Fatalf("missing 8")
	}
	v8 := b8[0]
	w8 := "abcdefg"
	p3 := newPair(diff(v8, v1, diff(v7, v1), diff(v4, v1)), diff(w8, w1, diff(w7, w1), diff(w4, w1)))

	fkey := []rune("abcdefg")
	variants := getVariants(p0, p1, p2, p3)
	matches := mappingNotUnique()
	for _, v := range variants {
		vm := []rune(v)
		remap := map[rune]rune{}
		for i := 0; i < len(fkey); i++ {
			remap[fkey[i]] = vm[i]
		}
		found := true
		fives, ok := byLen[5]
		if ok {
			for _, five := range fives {
				fiveMapped := replace(five, remap)
				_, ok := matches[fiveMapped]
				if !ok {
					found = false
					break
				}
			}
		}
		sixs, ok := byLen[6]
		if ok {
			for _, six := range sixs {
				sixMapped := replace(six, remap)
				_, ok := matches[sixMapped]
				if !ok {
					found = false
					break
				}
			}
		}
		if found {
			t := []int{0, 0, 0, 0}
			for i := 0; i < 4; i++ {
				ww := right[i]
				if len(ww) == 2 {
					t[i] = 1
				} else if len(ww) == 3 {
					t[i] = 7
				} else if len(ww) == 4 {
					t[i] = 4
				} else if len(ww) == 7 {
					t[i] = 8
				} else if len(ww) == 5 || len(ww) == 6 {
					vmapped := replace(ww, remap)
					vv, _ := matches[vmapped]
					t[i] = vv
				}
			}
			return t[0]*1000 + t[1]*100 + t[2]*10 + t[3]
		}
	}
	log.Fatalf("not found")

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	r1 := 0
	r2 := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		for i, part := range parts {
			parts[i] = sortString(part)
		}
		right := parts[len(parts)-4:]
		for _, v := range right {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				r1++
			}
		}
		left := parts[:len(parts)-5]
		r2 += decode(left, right)
	}
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)
}
