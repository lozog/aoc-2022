package main

import (
	"bufio"
	"fmt"
	"os"
)

// finds a shared element between two strings
func findSharedElement(first string, second string) rune {
	longerString := first
	shorterString := second
	if len(second) > len(first) {
		longerString = second
		shorterString = first
	}

	for _, firstElement := range shorterString {
		for _, secondElement := range longerString {
			if firstElement == secondElement {
				return firstElement
			}
		}
	}
	panic("could not find shared element")
}

// finds all shared elements between two strings
func findSharedElements(first string, second string) string {
	results := ""
	for _, firstElement := range first {
		for _, secondElement := range second {
			if firstElement == secondElement {
				results += string(firstElement)
			}
		}
	}
	return results
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

func p1() {
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

func p2() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	var elvesInGroup []string

	i := 0
	for fileScanner.Scan() {
		i += 1
		line := fileScanner.Text()
		// fmt.Println(line)

		elvesInGroup = append(elvesInGroup, line)

		if i%3 != 0 {
			// keep going until we get to each third elf
			continue
		}

		// find the element that is shared in first two elves
		sharedElementFirstTwo := findSharedElements(elvesInGroup[0], elvesInGroup[1])

		// find element shared between all three
		sharedElement := findSharedElement(sharedElementFirstTwo, elvesInGroup[2])

		// get that element's priority
		priority := runeToPriority(sharedElement)

		// add it to the sum
		sum += priority
		elvesInGroup = nil
	}
	readFile.Close()

	fmt.Printf("p2: %d\n", sum)
}

func main() {
	p1()
	p2()
}
