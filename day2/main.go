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

	var totalScore int

	for scanner.Scan() {
		var line = scanner.Text()
		var opponentHand = string(line[0])
		var ownHand = string(line[2])
		fmt.Printf("Opponent: %s vs Own: %s\n", opponentHand, ownHand)

		score := getScoreForRound(opponentHand, ownHand)
		fmt.Printf("Score: %d\n", score)
		totalScore += score
	}
	fmt.Printf("Total score: %d\n", totalScore)

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
		case "X": // lose
			return scissor + lost
		case "Y": // draw
			return rock + draw
		case "Z": // win
			return paper + won
		}
	case "B": // paper
		switch ownHand {
		case "X": // lose
			return rock + lost
		case "Y": // draw
			return paper + draw
		case "Z": // win
			return scissor + won
		}
	case "C": // scissor
		switch ownHand {
		case "X": // lose
			return paper + lost
		case "Y": // draw
			return scissor + draw
		case "Z": // win
			return rock + won
		}
	}
	return 0
}
