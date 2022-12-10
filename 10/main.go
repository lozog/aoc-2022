package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func incrementCycles(startCycle, cyclesToRun int, X int) (int, int) {
	cycle := startCycle
	signalStrength := 0
	for i := 0; i < cyclesToRun; i++ {
		// fmt.Printf("during cycle %d, X=%d\n", cycle, X)
		if (cycle+20)%40 == 0 {
			signalStrength = X * cycle
			// fmt.Printf("signal at cycle %d: %d\n", cycle, signalStrength)
		}
		cycle += 1
	}
	return cycle, signalStrength
}

func solution() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	X := 1
	cycle := 1
	sum := 0
	signalStrength := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		lineSplit := strings.Split(line, " ")
		instruction := lineSplit[0]
		var argument int
		if len(lineSplit) > 1 {
			argument, _ = strconv.Atoi(lineSplit[1])
		}

		if instruction == "noop" {
			cycle, signalStrength = incrementCycles(cycle, 1, X)
		} else if instruction == "addx" {
			cycle, signalStrength = incrementCycles(cycle, 2, X)
			X += argument
		}
		sum += signalStrength
	}
	readFile.Close()
	fmt.Printf("solution: %d\n", sum)
}

func main() {
	solution()
}
