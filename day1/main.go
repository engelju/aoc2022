package main

import (
	"bufio"
	"log"
	"os"
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

	var currentMaxCalories int64
	var overallMaxCalories int64

	for scanner.Scan() {
		println("current max: ", currentMaxCalories)
		var scannedText = scanner.Text()
		if scannedText == "" {
			println("found newline, resetting currentMaxCalories")
			if currentMaxCalories > overallMaxCalories {
				overallMaxCalories = currentMaxCalories
			}
			currentMaxCalories = 0
		} else {
			next, _ := strconv.ParseInt(scannedText, 0, 64)
			println("scanned value: ", next)
			currentMaxCalories += next
		}
		println("biggest max: ", overallMaxCalories)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
