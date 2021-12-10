package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func hasErrors(line string) (string, bool) {
	stack := []rune{}
	for _, ch := range []rune(line) {
		if ch == '(' || ch == '[' || ch == '{' || ch == '<' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' || ch == '>' {
			if len(stack) == 0 {
				return string(ch), true
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last == '(' && ch != ')' {
				return string(ch), true
			} else if last == '[' && ch != ']' {
				return string(ch), true
			} else if last == '<' && ch != '>' {
				return string(ch), true
			} else if last == '{' && ch != '}' {
				return string(ch), true
			}
		} else {
			log.Fatalf("invalid character %s", ch)
		}
	}
	return string(stack), false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	r1 := 0
	scores := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		ch, ok := hasErrors(line)
		if ok {
			if ch == ")" {
				r1 += 3
			} else if ch == "]" {
				r1 += 57
			} else if ch == "}" {
				r1 += 1197
			} else if ch == ">" {
				r1 += 25137
			}
		} else {
			s := 0
			for i := len(ch) - 1; i >= 0; i-- {
				x := ch[i]
				s *= 5
				if x == '(' {
					s += 1
				} else if x == '[' {
					s += 2
				} else if x == '{' {
					s += 3
				} else if x == '<' {
					s += 4
				}
			}
			scores = append(scores, s)
		}
	}
	fmt.Println("Part 1:", r1)
	sort.Ints(scores)
	r2 := scores[len(scores)/2]
	fmt.Println("Part 2:", r2)
}
