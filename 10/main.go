package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	X := 1
	cycle := 1
	sum := 0
	display := [][]rune{}
	// display = append(display, []rune{})
	curRow := -1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		lineSplit := strings.Split(line, " ")
		instruction := lineSplit[0]
		var argument int
		cyclesToRun := 1
		if len(lineSplit) > 1 {
			argument, _ = strconv.Atoi(lineSplit[1])
			cyclesToRun = 2
		}

		signalStrength := 0
		for i := 0; i < cyclesToRun; i++ {
			// fmt.Printf("signal at cycle %d: %d\n", cycle, X)

			// p1
			if (cycle+20)%40 == 0 {
				signalStrength = X * cycle
				// fmt.Printf("signal at cycle %d: %d\n", cycle, signalStrength)
			}

			curPixelPos := (cycle - 1) % 40
			if curPixelPos == 0 {
				display = append(display, []rune{})
				curRow += 1
			}

			elementToDraw := '.'
			if X-1 <= curPixelPos && curPixelPos <= X+1 {
				elementToDraw = '#'
			}

			// fmt.Printf("drawing %s for cycle %d\n", string(elementToDraw), cycle)
			display[curRow] = append(display[curRow], elementToDraw)

			cycle += 1
		}
		if instruction == "addx" {
			X += argument
		}
		sum += signalStrength
	}
	readFile.Close()
	fmt.Printf("p1: %d\n", sum)
	fmt.Print("p2:\n")
	for x := 0; x < len(display); x++ {
		for y := 0; y < len(display[x]); y++ {
			fmt.Printf("%s", string(display[x][y]))
		}
		fmt.Printf("\n")
	}
}

func main() {
	solution()
}
