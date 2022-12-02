package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func p1() {
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

	readFile, _ := os.Open("test.txt")
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

func findRequiredMoveForOutcome(requiredOutcome string, opponentMoveOutcomes map[string]string) string {
	for yourMove, outcome := range opponentMoveOutcomes {
		if outcome == requiredOutcome {
			return yourMove
		}
	}
	return "error"
}

func p2() {
	var opponentMoveMap = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	var requiredOutcomeMap = map[string]string{
		"X": "loss",
		"Y": "draw",
		"Z": "win",
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

	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		roundMoves := strings.Split(line, " ")

		opponentMove := opponentMoveMap[roundMoves[0]]

		requiredOutcome := requiredOutcomeMap[roundMoves[1]]
		// fmt.Printf("we need a %s against %s\n", requiredOutcome, opponentMove)
		// fmt.Printf("we need to find %s in %s\n", requiredOutcome, outcomeMap[opponentMove])
		yourMove := findRequiredMoveForOutcome(requiredOutcome, outcomeMap[opponentMove])

		outcome := outcomeMap[opponentMove][yourMove]

		// fmt.Printf("%s vs %s: %s\n", roundMoves[0], roundMoves[1], outcome)
		// fmt.Printf("moveScore: %d, outcomeScore: %d\n", moveScoreMap[yourMove], outcomeScoreMap[outcome])

		totalScore += moveScoreMap[yourMove] + outcomeScoreMap[outcome]
	}
	readFile.Close()
	// fmt.Printf("p2 final score: %d\n", totalScore)
}

func main() {
	p1()
	p2()
}
