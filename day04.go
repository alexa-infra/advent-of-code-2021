package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	boards []int
}

func (c *card) hasBingo(marked map[int]int) bool {
	for j := 0; j < 5; j++ {
		row := true
		for i := 0; i < 5; i++ {
			board := c.boards[j*5+i]
			_, ok := marked[board]
			if !ok {
				row = false
				break
			}
		}
		if row {
			return true
		}
	}
	for i := 0; i < 5; i++ {
		col := true
		for j := 0; j < 5; j++ {
			board := c.boards[j*5+i]
			_, ok := marked[board]
			if !ok {
				col = false
				break
			}
		}
		if col {
			return true
		}
	}
	return false
}

func (c *card) winScore(state map[int]int, n int) int {
	cc := 0
	for _, b := range c.boards {
		_, ok := state[b]
		if !ok {
			cc += b
		}
	}
	return cc * n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbersStr := strings.Split(scanner.Text(), ",")
	numbers := make([]int, 0, len(numbersStr))
	for _, num := range numbersStr {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}
	reNums := regexp.MustCompile(`\s+`)
	cards := []card{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			c := card{}
			for i := 0; i < 5; i++ {
				scanner.Scan()
				line = scanner.Text()
				parts := reNums.Split(strings.TrimSpace(line), -1)
				for _, num := range parts {
					n, _ := strconv.Atoi(num)
					c.boards = append(c.boards, n)
				}
			}
			cards = append(cards, c)
		}
	}
	r1 := 0
	state := map[int]int{}
	for _, n := range numbers {
		state[n] = 1
		found := false
		for _, c := range cards {
			if c.hasBingo(state) {
				found = true
				r1 = c.winScore(state, n)
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Println("Part 1:", r1)
	r2 := 0
	state = map[int]int{}
	for _, n := range numbers {
		state[n] = 1
		nextCards := []card{}
		for _, c := range cards {
			if c.hasBingo(state) {
				r2 = c.winScore(state, n)
			} else {
				nextCards = append(nextCards, c)
			}
		}
		if len(nextCards) == 0 {
			break
		}
		cards = nextCards
	}
	fmt.Println("Part 2:", r2)
}
