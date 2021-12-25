package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if parts[0] == "inp" {
			fmt.Printf("%s = inp[%d]\n", parts[1], i)
			i++
		} else if parts[0] == "add" {
			fmt.Printf("%s += %s\n", parts[1], parts[2])
		} else if parts[0] == "mul" {
			fmt.Printf("%s *= %s\n", parts[1], parts[2])
		} else if parts[0] == "div" {
			fmt.Printf("%s = %s // %s\n", parts[1], parts[1], parts[2])
		} else if parts[0] == "mod" {
			fmt.Printf("%s = %s %% %s\n", parts[1], parts[1], parts[2])
		} else if parts[0] == "eql" {
			fmt.Printf("%s = 1 if %s == %s else 0\n", parts[1], parts[1], parts[2])
		} else {
			fmt.Println("ERROR", parts[0])
		}
	}
}
