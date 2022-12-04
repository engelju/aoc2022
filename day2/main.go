package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var score int

	for scanner.Scan() {
		var line = scanner.Text()
		var opponentHand = string(line[0])
		var ownHand = string(line[2])
		fmt.Printf("Opponent: %s vs Own: %s\n", opponentHand, ownHand)

		score += getScoreForRound(opponentHand, ownHand)
		fmt.Printf("Score: %d\n", score)
	}
	fmt.Printf("Total score: %d\n", score)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var (
	rock    = 1
	paper   = 2
	scissor = 3
	won     = 6
	draw    = 3
	lost    = 0
)

func getScoreForRound(opponentHand, ownHand string) int {
	switch opponentHand {
	case "A": // rock
		switch ownHand {
		case "X": // rock
			return draw + rock
		case "Y": // paper
			return won + paper
		case "Z": // scissor
			return lost + scissor
		}
	case "B": // paper
		switch ownHand {
		case "X": // rock
			return lost + rock
		case "Y": // paper
			return draw + paper
		case "Z": // scissor
			return won + scissor
		}
	case "C": // scissor
		switch ownHand {
		case "X": // rock
			return won + rock
		case "Y": // paper
			return lost + paper
		case "Z": // scissor
			return draw + scissor
		}
	}
	return 0
}
