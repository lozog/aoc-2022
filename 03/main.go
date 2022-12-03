package main

import (
	"bufio"
	"fmt"
	"os"
)

// assumes len(first) == len(second)
func findSharedElement(first string, second string) rune {
	for _, firstElement := range first {
		for _, secondElement := range second {
			if firstElement == secondElement {
				return firstElement
			}
		}
	}
	panic("could not find shared element")
}

// a == 97
// A == 65
// runeToPriority("a") == 1
// runeToPriority("A") == 27
func runeToPriority(element rune) int {
	switch {
	case 97 <= element && element <= 122:
		// lowercase
		return int(element - 'a' + 1)
	case 65 <= element && element <= 90:
		// uppercase
		return int(element - 'A' + 26 + 1)
	default:
		panic("Rune out of expected range")
	}
}

func main() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		// split line into 2 halves
		first := line[0 : len(line)/2]
		second := line[len(line)/2:]

		// find the element that is shared in both
		sharedElement := findSharedElement(first, second)

		// get that element's priority
		priority := runeToPriority(sharedElement)

		// add it to the sum
		sum += priority
	}
	readFile.Close()

	fmt.Printf("p1: %d\n", sum)
}
