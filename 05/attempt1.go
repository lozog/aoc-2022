// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"

// 	stack "github.com/golang-collections/collections/stack"
// )

// func isLetter(a string) bool {
// 	return int(a[0]) >= int('A') && int(a[0]) <= int('Z')
// }

// // adapted from https://github.com/tinwatchman/2d-array-rotation/blob/master/dist/2d-array-rotation.js#L13
// func rotate90(a [][]string) [][]string {
// 	w := len(a)
// 	h := len(a[0])
// 	b := make([][]string, h)

// 	for y := 0; y < h; y++ {
// 		b[y] = make([]string, w)

// 		for x := 0; x < w; x++ {
// 			b[y][x] = a[w-1-x][y]
// 		}
// 	}

// 	return b
// }

// func p1() {
// 	readFile, _ := os.Open("test.txt")
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)

// 	stackInputData := [][]string{}
// 	stackMetaData := []int{}

// 	// scan each line of stack input into a 2d slice of strings
// 	for fileScanner.Scan() {
// 		line := fileScanner.Text()

// 		// check if we're on the line with the crate stack labels
// 		if string(line[1]) == "1" {
// 			stackMetaDataStrings := strings.Fields(line)
// 			for _, s := range stackMetaDataStrings {
// 				res, _ := strconv.Atoi(s)
// 				stackMetaData = append(stackMetaData, res)
// 			}
// 			break
// 		}

// 		stackInputData = append(stackInputData, strings.Split(line, ""))
// 	}

// 	// rotate 90 degrees clockwise
// 	rotated := rotate90(stackInputData)

// 	// create stacks
// 	stacks := []stack.Stack{}

// 	// read input into stacks
// 	for _, row := range rotated {
// 		first := row[0]

// 		// check if this row defines a stack
// 		if isLetter(first) {
// 			fmt.Println(row)

// 			// push crates onto stack
// 			for _, crate := range row {
// 				newStack := stack.Stack{}
// 				fmt.Printf("pushing %s\n", crate)
// 				newStack.Push(crate)
// 				stacks = append(stacks, newStack)
// 			}
// 		}
// 	}
// 	// fmt.Printf("%#v\n", stackInputData)

// 	fmt.Println(stacks[0].Pop())

// 	// fmt.Printf("p1 result: %d\n", sum)
// 	readFile.Close()
// }

// func main() {
// 	p1()
// }
