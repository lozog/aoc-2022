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

// https://stackoverflow.com/questions/35276022/unexpected-slice-append-behaviour
func makeFromSlice(sl []byte) []byte {
	result := make([]byte, len(sl))
	copy(result, sl)
	return result
}

func p1(maxCratesAtOnce int) {
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
		// fmt.Printf("need to move %d crate(s)\n", numCratesToMove)

		// do the moves
		for i := 0; i < numCratesToMove; i += maxCratesAtOnce {
			numCratesToMoveAtOnce := maxCratesAtOnce
			fmt.Printf("%d left to move\n", numCratesToMove-i)
			if numCratesToMove-i < maxCratesAtOnce {
				numCratesToMoveAtOnce = numCratesToMove - i
			}
			fmt.Printf("moving %d from %d to %d\n", numCratesToMoveAtOnce, sourceStackIdx, destStackIdx)
			fmt.Printf("moving %d crate(s)\n", numCratesToMoveAtOnce)
			sourceStack := stacks[sourceStackIdx]
			// fmt.Printf("source before: %s\n", stacks[sourceStackIdx])
			cratesToMove := sourceStack[0:numCratesToMoveAtOnce]
			// fmt.Printf("crates to move: %s\n", cratesToMove)
			stacks[sourceStackIdx] = sourceStack[numCratesToMoveAtOnce:]
			// fmt.Printf("source after: %s\n", stacks[sourceStackIdx])
			// fmt.Printf("dest before: %s\n", stacks[destStackIdx])
			// fmt.Printf("source after1: %s\n", stacks[sourceStackIdx])
			cratesToMove = append(makeFromSlice(cratesToMove), stacks[destStackIdx]...) // oh my god Golang why
			// fmt.Printf("source after2: %s\n", stacks[sourceStackIdx])
			stacks[destStackIdx] = cratesToMove
			// fmt.Printf("dest after: %s\n", stacks[destStackIdx])
			// fmt.Printf("source after3: %s\n", stacks[sourceStackIdx])
			fmt.Printf("%s\n", stacks)
		}

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
	p1(3)
}
