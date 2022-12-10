package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

// wtf Golang?!
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isAdjacent(a, b position) bool {
	return abs(a.x-b.x) <= 1 && abs(a.y-b.y) <= 1
}

func moveHead(start position, direction string) position {
	switch direction {
	case "R":
		return position{x: start.x + 1, y: start.y}
	case "L":
		return position{x: start.x - 1, y: start.y}
	case "U":
		return position{x: start.x, y: start.y + 1}
	case "D":
		return position{x: start.x, y: start.y - 1}
	}
	panic("invalid direction")
}

// given a tail that is not adjacent to the head, returns a new tail position
func moveTail(headPos position, tailPos position) position {
	// calculate direction to move tail in for each dimension
	xDirection := 0
	yDirection := 0
	if headPos.x > tailPos.x {
		xDirection = 1
	} else {
		xDirection = -1
	}

	if headPos.y > tailPos.y {
		yDirection = 1
	} else {
		yDirection = -1
	}

	// If the head is ever two steps directly up, down, left, or right from the tail,
	// the tail must also move one step in that direction so it remains close enough:
	if headPos.x == tailPos.x {
		return position{x: tailPos.x, y: tailPos.y + yDirection}
	}

	if headPos.y == tailPos.y {
		return position{x: tailPos.x + xDirection, y: tailPos.y}
	}

	// Otherwise, if the head and tail aren't touching and aren't in the same row or column,
	// the tail always moves one step diagonally to keep up:
	return position{x: tailPos.x + xDirection, y: tailPos.y + yDirection}
}

func solution() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tailVisitedPositions := make(map[string]int)
	headPos := position{x: 0, y: 0}
	tailPos := position{x: 0, y: 0}

	// must add starting position of tail
	positionAsStrIdx := strings.Join([]string{fmt.Sprint(tailPos.x), fmt.Sprint(tailPos.y)}, ",")
	tailVisitedPositions[positionAsStrIdx] += 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		lineSplit := strings.Split(line, " ")
		direction := lineSplit[0]
		distance, _ := strconv.Atoi(lineSplit[1])

		for i := 0; i < distance; i++ {
			// fmt.Printf("moving %s\n", direction)
			headPos = moveHead(headPos, direction)

			if !isAdjacent(headPos, tailPos) {
				tailPos = moveTail(headPos, tailPos)
				positionAsStrIdx := strings.Join([]string{fmt.Sprint(tailPos.x), fmt.Sprint(tailPos.y)}, ",")
				tailVisitedPositions[positionAsStrIdx] += 1
			}
			// fmt.Printf("headPos: %d, tailPos: %d\n", headPos, tailPos)
		}
	}
	readFile.Close()

	res := len(tailVisitedPositions)
	// fmt.Println(tailVisitedPositions)
	fmt.Printf("p1: %d\n", res)
}

func main() {
	solution()
}
