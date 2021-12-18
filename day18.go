package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type pair struct {
	cargo int
	left  *pair
	right *pair
}

func (p *pair) magnitude() int {
	if p.isRegular() {
		return p.cargo
	}
	return p.left.magnitude()*3 + p.right.magnitude()*2
}

func (p *pair) isRegular() bool {
	return p.left == nil || p.right == nil
}

func findSplit(head *pair) *pair {
	if head.isRegular() {
		if head.cargo >= 10 {
			return head
		}
	}
	if head.left != nil {
		ret := findSplit(head.left)
		if ret != nil {
			return ret
		}
	}
	if head.right != nil {
		ret := findSplit(head.right)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func doSplit(head, target *pair) {
	a := int(float64(target.cargo) / 2)
	b := int(float64(target.cargo)/2 + 0.5)
	target.cargo = -1
	target.left = &pair{a, nil, nil}
	target.right = &pair{b, nil, nil}
}

func findExplode(head *pair, level int) *pair {
	if level == 4 {
		if head.left != nil && head.left.isRegular() && head.right != nil && head.right.isRegular() {
			return head
		}
	}
	if head.left != nil {
		ret := findExplode(head.left, level+1)
		if ret != nil {
			return ret
		}
	}
	if head.right != nil {
		ret := findExplode(head.right, level+1)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func doExplode(head, target *pair) {
	flat := flatTree(head)
	for i, f := range flat {
		if f == target.left {
			if i > 0 {
				flat[i-1].cargo += target.left.cargo
			}
		}
		if f == target.right {
			if i+1 < len(flat) {
				flat[i+1].cargo += target.right.cargo
			}
		}
	}
	target.cargo = 0
	target.left = nil
	target.right = nil
}

func flatTree(head *pair) []*pair {
	if head.isRegular() {
		return []*pair{head}
	}
	ret := []*pair{}
	if head.left != nil {
		ret = append(ret, flatTree(head.left)...)
	}
	if head.right != nil {
		ret = append(ret, flatTree(head.right)...)
	}
	return ret
}

func nextClosed(text string) (int, bool) {
	cc := 0
	for i, ch := range text {
		if ch == '[' {
			cc++
		}
		if ch == ']' {
			cc--
			if cc == 0 {
				return i, true
			}
		}
	}
	return -1, false
}

func parse(text string) (*pair, int) {
	if text[0] == '[' {
		i, ok := nextClosed(text)
		if !ok {
			log.Fatalf("closed bracket not found")
		}
		left, lx := parse(text[1:i])
		if text[1+lx] != ',' {
			log.Fatalf("expected comma \"%s\"", text[1+lx:])
		}
		right, _ := parse(text[1+lx+1 : i])
		p := pair{-1, left, right}
		return &p, i + 1
	}
	if text[0] == ']' || text[0] == ',' {
		log.Fatalf("invalid symbol")
	}
	re := regexp.MustCompile(`^\d+`)
	match := re.FindStringIndex(text)
	if match == nil {
		log.Fatalf("expected number")
	}
	num, _ := strconv.Atoi(text[match[0]:match[1]])
	p := pair{num, nil, nil}
	return &p, match[1]
}

func reduce(head *pair) {
	for true {
		node := findExplode(head, 0)
		if node != nil {
			doExplode(head, node)
			continue
		}
		node = findSplit(head)
		if node != nil {
			doSplit(head, node)
			continue
		}
		break
	}
}

func add(a, b *pair) *pair {
	head := &pair{-1, a, b}
	reduce(head)
	return head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	var head *pair = nil
	for _, line := range lines {
		node, _ := parse(line)
		if head == nil {
			head = node
		} else {
			head = add(head, node)
		}
	}
	r1 := head.magnitude()
	fmt.Println("Part 1:", r1)
	r2 := 0
	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j {
				continue
			}
			node1, _ := parse(line1)
			node2, _ := parse(line2)
			sum := add(node1, node2)
			m := sum.magnitude()
			if m > r2 {
				r2 = m
			}
			node1, _ = parse(line1)
			node2, _ = parse(line2)
			sum = add(node2, node1)
			m = sum.magnitude()
			if m > r2 {
				r2 = m
			}
		}
	}
	fmt.Println("Part 2:", r2)
}
