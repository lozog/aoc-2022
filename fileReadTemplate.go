package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution() {
	readFile, _ := os.Open("test.txt")
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
