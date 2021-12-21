package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Player struct {
	pos   int
	score int
}

func (p Player) doMove(roll int) Player {
	p.pos = (p.pos-1+roll)%10 + 1
	p.score += p.pos
	return p
}

var diceOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

func diracStep(current, other Player, universes int) (currentWins, otherWins int) {
	for roll, times := range diceOutcomes {
		next := current.doMove(roll)
		if next.score >= 21 {
			currentWins += universes * times
		} else {
			otherWin, currentWin := diracStep(other, next, universes*times)
			currentWins += currentWin
			otherWins += otherWin
		}
	}
	return currentWins, otherWins
}

func main() {
	re := regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	pi1, pi2 := 0, 0
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		if match == nil {
			log.Fatalf("invalid format")
		}
		if match[1] == "1" {
			pi1, _ = strconv.Atoi(match[2])
		}
		if match[1] == "2" {
			pi2, _ = strconv.Atoi(match[2])
		}
	}
	p1, p2 := pi1, pi2
	ps1, ps2 := 0, 0
	n := 0
	r1 := 0
	for {
		move := (n+0)%100 + (n+1)%100 + (n+2)%100 + 3
		n += 3
		p1 = (p1-1+move)%10 + 1
		ps1 += p1
		if ps1 >= 1000 {
			r1 = ps2 * n
			break
		}
		move = (n+0)%100 + (n+1)%100 + (n+2)%100 + 3
		n += 3
		p2 = (p2-1+move)%10 + 1
		ps2 += p2
		if ps2 >= 1000 {
			r1 = ps1 * n
			break
		}
	}
	fmt.Println("Part 1:", r1)

	player1 := Player{pi1, 0}
	player2 := Player{pi2, 0}
	player1Wins, player2Wins := diracStep(player1, player2, 1)

	var winner int
	if player1Wins > player2Wins {
		winner = player1Wins
	} else {
		winner = player2Wins
	}
	fmt.Println("Part 2:", winner)
}
