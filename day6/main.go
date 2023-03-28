package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day6/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	//var min_marker = 0
	var maxMarker = 4
	var marker = make([]string, 4)

	for scanner.Scan() {
		var line = scanner.Text()
		//fmt.Printf("Line: %s\n", line)

		for i := 0; i < len(line); i++ {
			fmt.Printf("Character at %d Index Position = %c\n", i, line[i])
			if i < maxMarker {
				if contains(marker, string(line[i])) {
					fmt.Printf("already contained, shifting window\n")
					marker =
				} else {
					marker = append(marker, string(line[i]))
				}
			} else {
				fmt.Printf("markers: %s\n", marker)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
