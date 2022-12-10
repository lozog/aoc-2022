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

func solution(ropeLength int) {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tailVisitedPositions := make(map[string]int)
	TheRope := []position{}
	for i := 0; i < ropeLength; i++ {
		TheRope = append(TheRope, position{x: 0, y: 0})
	}
	tailIndex := ropeLength - 1

	// must add starting position of tail
	positionAsStrIdx := strings.Join([]string{fmt.Sprint(TheRope[tailIndex].x), fmt.Sprint(TheRope[tailIndex].y)}, ",")
	tailVisitedPositions[positionAsStrIdx] += 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		lineSplit := strings.Split(line, " ")
		direction := lineSplit[0]
		distance, _ := strconv.Atoi(lineSplit[1])

		// move head according to instructions
		for j := 0; j < distance; j++ {
			// fmt.Printf("moving %s\n", direction)
			TheRope[0] = moveHead(TheRope[0], direction)

			// propagate movement through rope
			for i := 0; i < tailIndex; i++ {
				if !isAdjacent(TheRope[i], TheRope[i+1]) {
					TheRope[i+1] = moveTail(TheRope[i], TheRope[i+1])

					// record any position the actual tail of the rope goes to
					if i+1 == tailIndex {
						positionAsStrIdx := strings.Join([]string{fmt.Sprint(TheRope[i+1].x), fmt.Sprint(TheRope[i+1].y)}, ",")
						tailVisitedPositions[positionAsStrIdx] += 1
					}
				}
			}
		}
	}
	readFile.Close()

	res := len(tailVisitedPositions)
	// fmt.Println(tailVisitedPositions)
	fmt.Printf("solution: %d\n", res)
}

func main() {
	solution(2)  // p1
	solution(10) // p2
}
