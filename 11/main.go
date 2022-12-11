package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// Monkey 0:
//
//	Starting items: 79, 98
//	Operation: new = old * 19
//	Test: divisible by 23
//	  If true: throw to monkey 2
//	  If false: throw to monkey 3
//
// ->
//
//	monkey {
//		id: 0,
//		items: [79, 98],
//		operator: '*',
//		operand: 19,
//		test: 23,
//		testTrueTarget: 2
//		testFalseTarget: 3
//	}
type monkey struct {
	id              int
	items           []int
	operator        string
	operand         string // i'm leaving this as a string for now, since it can also have the value "old"
	test            int
	testTrueTarget  int
	testFalseTarget int
	inspectionCount int
}

func strSliceToInt(a []string) []int {
	return funk.Map(a, func(x string) int {
		res, _ := strconv.Atoi(strings.TrimSpace(x))
		return res
	}).([]int)
}

func performOperation(old int, operator, operand string) int {
	var operandInt int
	if operand == "old" {
		operandInt = old
	} else {
		operandInt, _ = strconv.Atoi(operand)
	}

	if operator == "+" {
		return old + operandInt
	}
	return old * operandInt
}

func solution(relief bool) {
	readFile, _ := os.Open("data.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	monkeys := []monkey{}
	id := -1
	items := []int{}
	operator := ""
	operand := ""
	test := -1
	testTrueTarget := -1
	testFalseTarget := -1
	readMode := ""

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		switch readMode {
		case "":
			splitLine := strings.Split(line, " ")

			if splitLine[0] == "Monkey" {
				readMode = "items"
				id, _ = strconv.Atoi(string(splitLine[1][0]))
			}
		case "items":
			readMode = "operator"
			csv := line[18:]
			items = strSliceToInt(strings.Split(csv, ","))
		case "operator":
			readMode = "test"
			splitLine := strings.Split(strings.TrimSpace(line), " ")
			operator = splitLine[4]
			operand = splitLine[5]
		case "test":
			readMode = "testTrue"
			splitLine := strings.Split(strings.TrimSpace(line), " ")
			test, _ = strconv.Atoi(splitLine[3])
		case "testTrue":
			readMode = "testFalse"
			splitLine := strings.Split(strings.TrimSpace(line), " ")
			testTrueTarget, _ = strconv.Atoi(splitLine[5])
		case "testFalse":
			readMode = "done"
			splitLine := strings.Split(strings.TrimSpace(line), " ")
			testFalseTarget, _ = strconv.Atoi(splitLine[5])
		default:
			panic("unexpected readMode")
		}

		if readMode == "done" {
			readMode = ""
			monkey := monkey{
				id:              id,
				items:           items,
				operator:        operator,
				operand:         operand,
				test:            test,
				testTrueTarget:  testTrueTarget,
				testFalseTarget: testFalseTarget,
				inspectionCount: 0,
			}
			monkeys = append(monkeys, monkey)
		}
	}
	readFile.Close()

	LCM := funk.Reduce(monkeys, func(acc int, monkey monkey) int {
		return acc * monkey.test
	}, 1).(int)

	numRounds := 20
	if !relief {
		numRounds = 10000
	}

	for round := 0; round < numRounds; round++ {
		for _, monkey := range monkeys {
			// fmt.Println(monkey)
			// fmt.Printf("Monkey %d:\n", monkey.id)
			for _, item := range monkey.items {
				// fmt.Printf("  Monkey inspects an item with a worry level of %d\n", item)
				monkeys[monkey.id].inspectionCount += 1
				operationResult := performOperation(item, monkey.operator, monkey.operand)
				// fmt.Printf("    Worry level is changed to %d\n", operationResult)
				if relief {
					operationResult = operationResult / 3
					// fmt.Printf("    Worry level is divided by 3 to %d\n", operationResult)
				} else {
					operationResult %= LCM
				}
				testTarget := monkey.testFalseTarget
				if operationResult%monkey.test == 0 {
					testTarget = monkey.testTrueTarget
				}
				// fmt.Printf("    Item with worry level %d is thrown to monkey %d\n", operationResult, testTarget)
				monkeys[testTarget].items = append(monkeys[testTarget].items, operationResult)
			}
			monkeys[monkey.id].items = []int{}
		}
		// fmt.Printf("\nRound %d:\n", round)
		// for _, monkey := range monkeys {
		// 	fmt.Printf("Monkey %d: %d\n", monkey.id, monkey.items)
		// }
	}
	inspectionCounts := []int{}
	for _, monkey := range monkeys {
		inspectionCounts = append(inspectionCounts, monkey.inspectionCount)
	}
	sort.Ints(inspectionCounts)
	res := inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2]
	fmt.Printf("solution: %d\n", res)
}

func main() {
	solution(true)  // p1
	solution(false) // p2
}
