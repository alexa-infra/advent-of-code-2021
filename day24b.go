package main

import (
	"fmt"
	"sort"
)

var xa []int = []int{11, 12, 13, -5, -3, 14, 15, -16, 14, 15, -7, -11, -6, -11}
var xb []int = []int{16, 11, 12, 12, 12, 2, 11, 4, 12, 9, 10, 11, 6, 15}

func myFunc(inp []int) (z int) {
	// decompiled op-code program from the output of day24a
	z = 0
	for i := 0; i < 14; i++ {
		w := inp[i]
		z = myFuncStep(i, w, z)
	}
	return
}

func myFuncStep(i, w, z int) int {
	a := xa[i]
	b := xb[i]
	if a > 0 {
		return z*26 + w + b
	}
	zd := z % 26
	z = z / 26
	if zd != w-a {
		return z*26 + w + b
	}
	return z
}

func main() {
	zPrev := map[int]int{0: 1}
	steps := [][]int{
		[]int{0},
	}
	zLimit := 26 * 26 * 26 * 26
	for i := 0; i < 14; i++ {
		// for each iteration we build all possible outcomes of myFuncStep
		// while limiting possible z values, idk, it works
		zNext := map[int]int{}
		for z := range zPrev {
			if z > zLimit {
				continue
			}
			for w := 1; w <= 9; w++ {
				zz := myFuncStep(i, w, z)
				zNext[zz] = 1
			}
		}
		zNextArr := []int{}
		for z := range zNext {
			zNextArr = append(zNextArr, z)
		}
		steps = append(steps, zNextArr)
		zPrev = zNext
	}
	found := false
	for _, z := range steps[14] {
		if z == 0 {
			// 14th iteration outcomes should contain zero
			// this is the MONAD-valid condition
			found = true
			break
		}
	}
	if !found {
		// if we haven't found it, then increase more zLimit :)
		fmt.Println("z=0 not found, need more zLimit!")
		return
	}
	steps[14] = []int{0}
	type pair struct {
		x, y int
	}
	revSteps := [][]int{}
	revPrev := map[pair]int{pair{0, 0}: 0}
	for i := 14; i >= 1; i-- {
		// now we rebuild which (w, z)-pairs have led to the valid
		// condition, so we use cached z-values from the previous step
		// and go in the oposite direction
		revNext := map[pair]int{}
		for _, z := range steps[i-1] {
			for w := 1; w <= 9; w++ {
				zz := myFuncStep(i-1, w, z)
				for p := range revPrev {
					if p.x == zz {
						pp := pair{z, w}
						revNext[pp] = 1
						break
					}
				}
			}
		}
		// one each step we need only w-values
		// so we can significantly decrease bruteforce
		revStep := []int{}
		hash := map[int]int{}
		for p := range revNext {
			_, ok := hash[p.y]
			if !ok {
				hash[p.y] = 1
				revStep = append(revStep, p.y)
			}
		}
		revSteps = append(revSteps, revStep)
		revPrev = revNext
	}
	// each step is the position in the number
	// for each step we have a list of valid numbers
	// now we do the bruteforce :)
	for _, s := range revSteps {
		// since we only need the largest number
		// we sort each list from bigger to lower
		// so the very first match would be the biggest
		// possible MONAD-value
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
	}
	res := make([]int, 14)
	found = false
	var req func(pos int)
	fmt.Printf("Part 1: ")
	req = func(pos int) {
		if found {
			return
		}
		if pos == 14 {
			z := myFunc(res)
			if z == 0 {
				// we only need the very first number
				for _, w := range res {
					fmt.Printf("%d", w)
				}
				found = true
			}
			return
		}
		for _, w := range revSteps[pos] {
			if found {
				break
			}
			res[13-pos] = w
			req(pos + 1)
		}
	}
	req(0)
	fmt.Println()
	for _, s := range revSteps {
		// and for the second part we need the lowest number
		// so just do the sort from lowest to biggest
		// and take the first occurance
		sort.Sort(sort.IntSlice(s))
	}
	found = false
	fmt.Printf("Part 2: ")
	req(0)
	fmt.Println()
}
