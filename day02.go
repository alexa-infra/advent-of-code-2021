package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	dir  string
	move int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []command{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		move, _ := strconv.Atoi(parts[1])
		lines = append(lines, command{dir, move})
	}
	posX, posY := 0, 0
	for _, line := range lines {
		if line.dir == "forward" {
			posX += line.move
		} else if line.dir == "up" {
			posY -= line.move
		} else if line.dir == "down" {
			posY += line.move
		} else {
			log.Fatalf("wrong direction")
		}
	}
	r1 := posX * posY
	fmt.Println("Part 1:", r1)
	posX, posY = 0, 0
	aim := 0
	for _, line := range lines {
		if line.dir == "forward" {
			posX += line.move
			posY += aim * line.move
		} else if line.dir == "up" {
			aim -= line.move
		} else if line.dir == "down" {
			aim += line.move
		} else {
			log.Fatalf("wrong direction")
		}
	}
	r2 := posX * posY
	fmt.Println("Part 2:", r2)
}
