package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func changeDirectory(path []string, dstDir string) []string {

	if dstDir == ".." {
		if len(path) == 1 {
			// if root, stay on the root
			return path
		}
		// handle moving up a folder
		return path[:len(path)-1]
	}

	return append(path, dstDir)
}

// [/ foo bar baz] => "/foo/bar/baz"
func dirSliceToString(path []string) string {
	if len(path) == 1 {
		return path[0]
	}
	res := strings.Join(path[1:], "/")
	return path[0] + res
}

func solution() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// curDir := ""
	curDir := []string{}
	curCmd := ""
	dirSizes := make(map[string]int)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Printf("%s\n", line)
		splitLine := strings.Split(line, " ")

		if splitLine[0] == "$" {
			curCmd = splitLine[1]

			if curCmd == "cd" {
				dstDir := splitLine[2]
				curDir = changeDirectory(curDir, dstDir)
				// fmt.Printf("changing to new dir: %s\n", curDir)
			}
		} else if curCmd == "ls" {
			if splitLine[0] != "dir" {
				// record file sizes

				fileSize, _ := strconv.Atoi(splitLine[0])

				// also add filesize to this folder and all its parents
				dir := curDir
				for dirSliceToString(dir) != "/" {
					dirSizes[dirSliceToString(dir)] += fileSize
					dir = changeDirectory(dir, "..")
				}
				// finally, add it to the root
				dirSizes["/"] += fileSize

			}
		}
	}
	readFile.Close()

	// fmt.Printf("%+v\n", dirSizes)

	// sum := 0
	// for _, dirSize := range dirSizes {
	// 	if dirSize <= 100000 {
	// 		sum += dirSize
	// 	}
	// }
	// fmt.Printf("p1: %d\n", sum)

	curUnusedSpace := 70000000 - dirSizes["/"]
	neededUnusedSpace := 30000000 - curUnusedSpace // assumes 30000000 > curUnusedSpace

	curMin := ""
	curMinSize := 0
	for path, dirSize := range dirSizes {
		if dirSize >= neededUnusedSpace && (curMin == "" || dirSize < curMinSize) {
			curMin = path
			curMinSize = dirSize
		}
	}
	fmt.Printf("p2: %s\n", curMin)
	fmt.Printf("p2: %d\n", dirSizes[curMin])

}

func main() {
	solution()
}
