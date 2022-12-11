package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stacks [9][]string
	var secondPhase bool

	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Printf("Line: %s\n", line)

		if line == "" {
			continue
		}

		if !secondPhase {
			if unicode.IsNumber(rune(line[1])) || unicode.IsNumber(rune(line[5])) || unicode.IsNumber(rune(line[9])) {
				secondPhase = true
				for i := range stacks {
					stacks[i] = reverseStack(stacks[i])
				}
				continue
			}

			startIdx := 1
			shiftIdx := 4

			for i := range stacks {
				if string(line[startIdx]) != " " {
					stacks[i] = append(stacks[i], string(line[startIdx]))
				}
				startIdx = startIdx + shiftIdx
			}

		} else {
			words := strings.Split(line, " ")

			quantity, _ := strconv.Atoi(words[1])
			from, _ := strconv.Atoi(words[3])
			to, _ := strconv.Atoi(words[5])

			fmt.Printf("move %d from %d to %d\n", quantity, from, to)
			stacks[from-1], stacks[to-1] = move(stacks[from-1], stacks[to-1], quantity)
		}

		for k, v := range stacks {
			fmt.Printf("Stack %d: %s\n", k, v)
		}

		fmt.Println("---")
	}

	for _, stack := range stacks {
		fmt.Printf("Top elem in Stack is: %s\n", stack[len(stack)-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func reverseStack(oldStack []string) []string {
	var newStack []string
	for len(oldStack) > 0 {
		topElemIdx := len(oldStack) - 1
		newStack = append(newStack, oldStack[topElemIdx])
		oldStack = oldStack[:topElemIdx]
	}
	return newStack
}

func move(fromStack []string, toStack []string, howMany int) ([]string, []string) {
	for howMany > 0 {
		topElemIdx := len(fromStack) - 1
		toStack = append(toStack, fromStack[topElemIdx])
		fromStack = fromStack[:topElemIdx]
		howMany--
	}

	return fromStack, toStack
}
