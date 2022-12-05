package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var totalOverlappingPairs int

	for scanner.Scan() {
		var line = scanner.Text()
		splitLine := strings.Split(line, ",")

		var elf1 []int
		for _, i := range strings.Split(splitLine[0], "-") {
			c, _ := strconv.Atoi(i)
			elf1 = append(elf1, c)
		}
		fmt.Printf("Elf 1: %d\n", elf1)

		var elf2 []int
		for _, i := range strings.Split(splitLine[1], "-") {
			c, _ := strconv.Atoi(i)
			elf2 = append(elf2, c)
		}
		fmt.Printf("Elf 2: %d\n", elf2)

		// check if elf2 is contained in elf1
		if elf2[0] >= elf1[0] && elf2[0] <= elf1[1] {
			fmt.Printf("--> Elf1: %d contains Elf2: %d\n", elf1, elf2)
			totalOverlappingPairs++
			continue
		}

		// check if elf1 is contained in elf2
		if elf1[0] >= elf2[0] && elf1[0] <= elf2[1] {
			fmt.Printf("--> Elf2: %d contains Elf1: %d\n", elf2, elf1)
			totalOverlappingPairs++
			continue
		}

		fmt.Println("---")
	}
	fmt.Printf("Total overlapping pairs: %d\n", totalOverlappingPairs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
