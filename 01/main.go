package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	curMax := 0
	curElf := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		cur, err := strconv.Atoi(line)
		if err != nil {
			// fmt.Println("end of elf")
			if curElf > curMax {
				curMax = curElf
			}
			curElf = 0
			continue
		}
		curElf += cur
	}

	readFile.Close()

	fmt.Printf("res: %d\n", curMax)
}
