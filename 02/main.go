package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var opponentMoveMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var yourMoveMap = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

// maps opponent moves to an array of outcomes of your moves
// e.g. outcomeMap["paper"]["scissors"] == win because scissors beats paper
var outcomeMap = map[string]map[string]string{ // what the hell, Golang
	"rock": {
		"rock":     "draw",
		"paper":    "win",
		"scissors": "loss",
	},
	"paper": {
		"rock":     "loss",
		"paper":    "draw",
		"scissors": "win",
	},
	"scissors": {
		"rock":     "win",
		"paper":    "loss",
		"scissors": "draw",
	},
}

var moveScoreMap = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var outcomeScoreMap = map[string]int{
	"loss": 0,
	"draw": 3,
	"win":  6,
}

func main() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		roundMoves := strings.Split(line, " ")

		opponentMove := opponentMoveMap[roundMoves[0]]
		yourMove := yourMoveMap[roundMoves[1]]

		outcome := outcomeMap[opponentMove][yourMove]

		// fmt.Printf("%s vs %s: %s\n", roundMoves[0], roundMoves[1], outcome)
		// fmt.Printf("moveScore: %d, outcomeScore: %d\n", moveScoreMap[yourMove], outcomeScoreMap[outcome])

		totalScore += moveScoreMap[yourMove] + outcomeScoreMap[outcome]
	}
	readFile.Close()
	fmt.Printf("p1 final score: %d\n", totalScore)
}
