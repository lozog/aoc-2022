package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLetter(a byte) bool {
	return int(a) >= int('A') && int(a) <= int('Z')
}

func p1() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	stacks := [][]byte{}

	// scan each line of stack input into a 2d slice of strings
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			break
		}

		numStacks := (len(line) + 1) / 4
		if len(stacks) == 0 {
			for i := 0; i < numStacks; i++ {
				stacks = append(stacks, []byte{})
			}
		}

		for i := 0; i < numStacks; i++ {
			crateCandidate := line[(i*4)+1]

			// check if it is a crate
			if isLetter(crateCandidate) {
				stacks[i] = append(stacks[i], crateCandidate)
			}
		}
	}

	fmt.Printf("%s\n", stacks)

	// now we have our stacks

	// time to read + follow the instructions

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " ")

		numCratesToMove, _ := strconv.Atoi(split[1])
		sourceStackIdx, _ := strconv.Atoi(split[3])
		destStackIdx, _ := strconv.Atoi(split[5])

		// stacks are counted in inputs starting at 1, but we'll be using them as slice indices
		sourceStackIdx -= 1
		destStackIdx -= 1

		// do the moves
		for i := 0; i < numCratesToMove; i++ {
			sourceStack := &stacks[sourceStackIdx]
			// fmt.Printf("source before: %s\n", *sourceStack)
			crateToMove := (*sourceStack)[0]
			*sourceStack = (*sourceStack)[1:]
			// fmt.Printf("source after: %s\n", *sourceStack)
			destStack := &stacks[destStackIdx]
			// fmt.Printf("dest before: %s\n", *destStack)
			*destStack = append([]byte{crateToMove}, *destStack...)
			// fmt.Printf("dest after: %s\n", *destStack)
			// fmt.Println()
		}
		// fmt.Println()

	}
	readFile.Close()

	// get answers
	res := ""
	for i := 0; i < len(stacks); i++ {
		res += string(stacks[i][0])
	}
	fmt.Printf("%s\n", stacks)

	fmt.Printf("p1 result: %s\n", res)
}

func main() {
	p1()
}
