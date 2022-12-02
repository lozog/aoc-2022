package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := make([]int, 0)
	curElf := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		cur, err := strconv.Atoi(line)
		if err != nil {
			// fmt.Println("end of elf")
			elves = append(elves, curElf)
			curElf = 0
			continue
		}
		curElf += cur
	}
	elves = append(elves, curElf) // don't forget to append the last elf

	readFile.Close()

	sort.Ints(elves)

	fmt.Printf("p1: %d\n", elves[len(elves)-1])

	p2_result := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Printf("p2: %d\n", p2_result)
}
