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
	folderSizes := make(map[string]int)

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
			if splitLine[0] == "dir" {
				// curDir = changeDirectory(curDir, splitLine[1])
				// curDirAsString := dirSliceToString(curDir)
				// folderSizes[curDirAsString] = 0
				// fmt.Printf("changing to new dir: %s\n", curDir)
			} else {
				// record file sizes
				res, _ := strconv.Atoi(splitLine[0])
				// curDirAsString := dirSliceToString(curDir)
				// folderSizes[curDirAsString] += res

				// also add filesize to every parent folder
				dir := curDir
				for dirSliceToString(dir) != "/" {
					folderSizes[dirSliceToString(dir)] += res
					// fmt.Printf("adding %d to %s: %d\n", res, dirSliceToString(dir), folderSizes[dirSliceToString(dir)])
					dir = changeDirectory(dir, "..")
				}
				// finally, add it to the root
				folderSizes["/"] += res
				// fmt.Printf("adding %d to %s: %d\n", res, "/", folderSizes["/"])

			}
		}
	}
	readFile.Close()

	// fmt.Printf("%+v\n", folderSizes)

	// sum := 0
	// for _, folderSize := range folderSizes {
	// 	if folderSize <= 100000 {
	// 		sum += folderSize
	// 	}
	// }
	// fmt.Printf("p1: %d\n", sum)

	curUnusedSpace := 70000000 - folderSizes["/"]
	neededUnusedSpace := 30000000 - curUnusedSpace // assumes 30000000 > curUnusedSpace

	curMin := ""
	curMinSize := 0
	for path, folderSize := range folderSizes {
		if folderSize >= neededUnusedSpace && (curMin == "" || folderSize < curMinSize) {
			curMin = path
			curMinSize = folderSize
		}
	}
	fmt.Printf("p2: %s\n", curMin)
	fmt.Printf("p2: %d\n", folderSizes[curMin])

}

func main() {
	solution()
}
