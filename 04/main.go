package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// converts string array to int array
func stringSliceToInt(in []string) []int {
	var out = make([]int, len(in))

	for idx, i := range in {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		out[idx] = j
	}
	return out
}

func p1() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		pairs := strings.Split(line, ",")

		pair1 := stringSliceToInt(strings.Split(pairs[0], "-"))
		pair2 := stringSliceToInt(strings.Split(pairs[1], "-"))

		// fmt.Println(line)

		if pair1[0] <= pair2[0] && pair2[1] <= pair1[1] {
			// check if pair2 is within pair1's range
			sum += 1
		} else if pair2[0] <= pair1[0] && pair1[1] <= pair2[1] {
			// check if pair1 is within pair2's range
			sum += 1
		}
	}
	readFile.Close()

	fmt.Printf("p1 result: %d\n", sum)
}

func p2() {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		pairs := strings.Split(line, ",")

		pair1 := stringSliceToInt(strings.Split(pairs[0], "-"))
		pair2 := stringSliceToInt(strings.Split(pairs[1], "-"))

		// fmt.Println(line)

		if pair1[0] <= pair2[0] && pair2[0] <= pair1[1] {
			// if pair1 comes first and pair2 starts before pair1 ends
			// fmt.Printf("%d <= %d && %d <= %d\n", pair1[0], pair2[0], pair2[0], pair1[1])
			sum += 1
		} else if pair2[0] <= pair1[0] && pair1[0] <= pair2[1] {
			// if pair2 comes first and pair1 starts before pair2 ends
			// fmt.Printf("%d <= %d && %d <= %d\n", pair2[0], pair1[0], pair1[0], pair2[1])
			sum += 1
		}
	}
	readFile.Close()

	fmt.Printf("p2 result: %d\n", sum)
}

func main() {
	// p1()
	p2()
}
