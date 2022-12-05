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
	var counter int
	var triplet [3]string

	for scanner.Scan() {
		var line = scanner.Text()

		triplet[counter] = line
		counter++

		if counter == 3 {
			var commonLetters = compareCompartments(triplet[0], triplet[1], triplet[2])
			for letter := range commonLetters {
				fmt.Printf("Priority of common letter: %s is: %d\n", letter, commonLetters[letter])
				totalPriorities += commonLetters[letter]
			}
			counter = 0
		}
	}
	fmt.Printf("Total score: %d\n", totalPriorities)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compareCompartments(string1, string2, string3 string) map[string]int {

	commonLetters := make(map[string]int)

	for i := 0; i < len(string1); i++ {
		for j := 0; j < len(string2); j++ {
			for k := 0; k < len(string3); k++ {
				char1 := string(string1[i])
				char2 := string(string2[j])
				char3 := string(string3[k])
				if (char1 == char2) && (char1 == char3) {
					commonLetters[char3] = calculatePriority(char3)
					continue
				}
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
