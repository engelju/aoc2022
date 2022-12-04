package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day3/input.txt")
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
		for letter := range commonLetters {
			fmt.Printf("Priority of common letter: %s is: %d\n", letter, commonLetters[letter])
			totalPriorities += commonLetters[letter]
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Total score: %d\n", totalPriorities)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compareCompartments(string1, string2 string) map[string]int {

	commonLetters := make(map[string]int)

	for i := 0; i < len(string1); i++ {
		for j := 0; j < len(string2); j++ {
			char1 := string(string1[i])
			char2 := string(string2[j])
			if char1 == char2 {
				commonLetters[char2] = calculatePriority(char2)
				continue
			}
		}
	}

	return commonLetters
}

func calculatePriority(charAsString string) int {
	runeArray := []rune(charAsString)
	char := runeArray[0]
	if char >= 'a' && char <= 'z' {
		return int(char - 96)
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 64 + 26)
	}
	return 0
}
