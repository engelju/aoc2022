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
	f, err := os.Open("day1/test.txt")
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
		fmt.Println("current max: ", currentMaxCalories)
		if line == "" {
			fmt.Println("found newline, resetting currentMaxCalories")
			elves = append(elves, currentMaxCalories)
			currentMaxCalories = 0
		} else {
			next, _ := strconv.Atoi(line)
			fmt.Println("scanned value: ", next)
			currentMaxCalories += next
		}
	}

	fmt.Println(elves)
	sort.Ints(elves)
	fmt.Println(elves)

	fmt.Println(elves[len(elves)-1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
