package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func solution() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	treeHeightMap := [][]int{}
	treeVisibilityMap := [][]bool{}

	row := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		lineSlice := funk.Map(strings.Split(line, ""), func(x string) int {
			res, _ := strconv.Atoi(x)
			return res
		}).([]int)

		boolSlice := funk.Map(lineSlice, func(_ int) bool {
			return false
		}).([]bool)

		// fmt.Printf("%d\n", lineSlice)

		treeHeightMap = append(treeHeightMap, lineSlice)
		treeVisibilityMap = append(treeVisibilityMap, boolSlice)
		// fmt.Printf("%d\n", treeHeightMap)

		curMaxHeight := -1
		// for each row, check which trees are visible from the left
		for col, treeHeight := range lineSlice {
			if treeHeight > curMaxHeight {
				curMaxHeight = treeHeight
				treeVisibilityMap[row][col] = true
			}
		}

		// now the right
		curMaxHeight = -1
		for col := len(lineSlice) - 1; col >= 0; col-- {
			if lineSlice[col] > curMaxHeight {
				curMaxHeight = lineSlice[col]
				treeVisibilityMap[row][col] = true
			}
		}
		row += 1
	}
	readFile.Close()

	// for row := 0; row < len(treeHeightMap); row++ {
	// 	fmt.Printf("%t\n", treeVisibilityMap[row])
	// }

	for col := 0; col < len(treeHeightMap[0]); col++ {
		// for each column, check which trees are visible from the top
		curMaxHeight := -1
		for row := 0; row < len(treeHeightMap); row++ {
			if treeHeightMap[row][col] > curMaxHeight {
				curMaxHeight = treeHeightMap[row][col]
				treeVisibilityMap[row][col] = true
			}
		}

		// now, from the bottom
		curMaxHeight = -1
		for row := len(treeHeightMap) - 1; row >= 0; row-- {
			if treeHeightMap[row][col] > curMaxHeight {
				curMaxHeight = treeHeightMap[row][col]
				treeVisibilityMap[row][col] = true
			}
		}
	}

	sum := 0
	for row := 0; row < len(treeHeightMap); row++ {
		// fmt.Printf("%t\n", treeVisibilityMap[row])
		for col := 0; col < len(treeHeightMap[0]); col++ {
			if treeVisibilityMap[row][col] {
				sum += 1
			}
		}
	}
	fmt.Printf("p1: %d\n", sum)
}

func main() {
	solution()
}
