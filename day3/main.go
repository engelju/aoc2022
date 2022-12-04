package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day3/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var totalPriorities int

	for scanner.Scan() {
		var line = scanner.Text()

		split := len(line) / 2
		compartment1 := line[:split]
		compartment2 := line[split:]
		fmt.Printf("Compartment 1: %s and Compartment 2: %s\n", compartment1, compartment2)

		var commonLetters = compareCompartments(compartment1, compartment2)
		fmt.Printf("Common letters in both compartments:")
		for _, letter := range commonLetters {
			fmt.Print(letter)
		}
		fmt.Printf("\n")

		//var opponentHand = string(line[0])
		//var ownHand = string(line[2])
		//fmt.Printf("Opponent: %s vs Own: %s\n", opponentHand, ownHand)

		//score := getScoreForRound(opponentHand, ownHand)
		//fmt.Printf("Score: %d\n", score)
		//totalPriorities += score
	}
	fmt.Printf("Total score: %d\n", totalPriorities)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compareCompartments(string1, string2 string) []string {

	var commonLetters []string

	//var map_1 map[string]bool

	//visitedURL := map[string]bool{
	//	"http://www.google.com": true,
	//	"https://paypal.com":    true,
	//}
	//
	//if visitedURL[thisSite] {
	//	fmt.Println("Already been here.")
	//}

	for i := 0; i < len(string1); i++ {
		for j := 0; j < len(string2); j++ {
			char1 := string(string1[i])
			char2 := string(string2[j])
			if char1 == char2 {
				//map_1[char2] = true

				commonLetters = append(commonLetters, char2)
				continue
			}
		}
	}

	return commonLetters
}
