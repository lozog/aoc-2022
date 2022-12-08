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

func solution(windowLength int) {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		for i := range line {
			// fmt.Println(i)
			endOfWindow := min(len(line)-1, i+windowLength)
			window := line[i:endOfWindow]
			// fmt.Println(window)
			if allUniqueChars(window) {
				fmt.Printf("result: %d\n", endOfWindow)
				break
			}

		}
	}
	readFile.Close()
}

func main() {
	solution(4)  // p1
	solution(14) // p2
}
