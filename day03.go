package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func filter(arr []string, p int, most bool) []string {
	ones := 0
	for _, el := range arr {
		if rune(el[p]) == '1' {
			ones++
		}
	}
	zeros := len(arr) - ones
	keep := 'x'
	n := 0
	if most {
		if ones > zeros {
			keep = '1'
			n = ones
		} else if zeros > ones {
			keep = '0'
			n = zeros
		} else {
			keep = '1'
			n = ones
		}
	} else {
		if ones > zeros {
			keep = '0'
			n = zeros
		} else if zeros > ones {
			keep = '1'
			n = ones
		} else {
			keep = '0'
			n = zeros
		}
	}
	result := make([]string, 0, n)
	for _, el := range arr {
		if rune(el[p]) == keep {
			result = append(result, el)
		}
	}
	return result
}

func filterTerm(arr []string, most bool) string {
	i := 0
	for len(arr) > 1 {
		arr = filter(arr, i, most)
		i++
	}
	return arr[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	n := len(lines)
	ones := make([]int, 0, len(lines[0]))
	for j, line := range lines {
		for i, ch := range []rune(line) {
			if ch == '1' {
				if j == 0 {
					ones = append(ones, 1)
				} else {
					ones[i]++
				}
			} else if ch == '0' {
				if j == 0 {
					ones = append(ones, 0)
				}
			} else {
				log.Fatalf("invalid character")
			}
		}
	}
	gamma := make([]rune, 0, n)
	epsilon := make([]rune, 0, n)
	for _, val := range ones {
		zeros := n - val
		if val > zeros {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}
	gammaStr := string(gamma)
	epsilonStr := string(epsilon)
	gammaRate, _ := new(big.Int).SetString(gammaStr, 2)
	epsilonRate, _ := new(big.Int).SetString(epsilonStr, 2)
	r1 := new(big.Int).Mul(gammaRate, epsilonRate)
	fmt.Println("Part 1:", r1.Text(10))

	oxygenStr := filterTerm(lines, true)
	co2Str := filterTerm(lines, false)
	oxygenRate, _ := new(big.Int).SetString(oxygenStr, 2)
	co2Rate, _ := new(big.Int).SetString(co2Str, 2)
	r2 := new(big.Int).Mul(oxygenRate, co2Rate)
	fmt.Println("Part 2:", r2.Text(10))
}
