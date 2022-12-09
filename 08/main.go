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
	// 	fmt.Printf("%d\n", treeHeightMap[row])
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

	height := len(treeHeightMap)
	width := len(treeHeightMap[0])

	sum := 0
	for row := 0; row < height; row++ {
		// fmt.Printf("%t\n", treeVisibilityMap[row])
		for col := 0; col < width; col++ {
			if treeVisibilityMap[row][col] {
				sum += 1
			}
		}
	}
	fmt.Printf("p1: %d\n", sum)

	// part 2:
	// from every tree, look in every direction and calculate its scenic score
	maxScenicScore := 0
	for row := 0; row < height; row++ {
		// fmt.Printf("%t\n", treeVisibilityMap[row])
		for col := 0; col < width; col++ {
			scenicScoreParts := [4]int{0, 0, 0, 0}

			// look right
			for y := col + 1; y < width; y++ {
				scenicScoreParts[0] += 1
				if treeHeightMap[row][y] >= treeHeightMap[row][col] {
					break
				}
			}

			// look left
			for y := col - 1; y >= 0; y-- {
				scenicScoreParts[1] += 1
				if treeHeightMap[row][y] >= treeHeightMap[row][col] {
					break
				}
			}

			// look down
			for x := row + 1; x < height; x++ {
				scenicScoreParts[2] += 1
				if treeHeightMap[x][col] >= treeHeightMap[row][col] {
					break
				}
			}

			// look up
			for x := row - 1; x >= 0; x-- {
				scenicScoreParts[3] += 1
				if treeHeightMap[x][col] >= treeHeightMap[row][col] {
					break
				}
			}

			// fmt.Println(scenicScoreParts)

			scenicScore := funk.Reduce(scenicScoreParts, '*', 1).(int)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	fmt.Printf("p2: %d\n", maxScenicScore)
}

func main() {
	solution()
}
