package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// what the hell, Golang?
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func allUniqueChars(s string) bool {
	tmp := ""
	for _, char := range s {
		if strings.ContainsRune(tmp, char) {
			return false
		}
		tmp += string(char)
	}
	return true
}

func p1() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		for i := range line {
			// fmt.Println(i)
			endOfWindow := min(len(line)-1, i+4)
			window := line[i:endOfWindow]
			// fmt.Println(window)
			if allUniqueChars(window) {
				fmt.Printf("p1 result: %d\n", endOfWindow)
				break
			}

		}
	}
}

func main() {
	p1()
}
