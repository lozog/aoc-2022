package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(windowLength int) {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
	}
	readFile.Close()
}

func main() {
	solution()
}
