package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var currentMaxCalories int
	var elves []int

	for scanner.Scan() {
		var line = scanner.Text()
		if line != "" {
			next, _ := strconv.Atoi(line)
			currentMaxCalories += next
			continue
		}
		elves = append(elves, currentMaxCalories)
		currentMaxCalories = 0
	}

	// add the last one manually
	elves = append(elves, currentMaxCalories)

	sort.Ints(elves)

	// add up the top 3 elves
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
