package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for _, line := range readFromFile("input.txt") {
		var characters []string
		fmt.Printf("Line: %s\n", line)
		for i := 0; i < len(line); i++ {
			// fmt.Printf("Character at %d Index Position = %c\n", i, line[i])

			if len(characters) < 4 {
				characters = append(characters, string(line[i]))
			} else if len(characters) == 4 {
				if charsAreDifferent(characters) {
					fmt.Printf("Found four different characters: %s\n", characters)
					fmt.Printf("Index: %d\n", i)
					break
				} else {
					characters = characters[1:]
					characters = append(characters, string(line[i]))
				}
			}
			// fmt.Printf("characters: %s\n\n", characters)
		}
	}
}

func readFromFile(filename string) []string {
	var lines []string

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func charsAreDifferent(characters []string) bool {
	var allDifferent = true
	for i := 0; i < len(characters); i++ {
		for j := i + 1; j < len(characters); j++ {
			if characters[i] == characters[j] {
				// fmt.Printf("Characters %s and %s are the same\n", characters[i], characters[j])
				allDifferent = false
			}
		}
	}
	return allDifferent
}
