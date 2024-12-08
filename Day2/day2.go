package main

// Input is a list of reports, which look like ints
// These reports are safe if:
//	- The levels are either all increasing or all decreasing, AND
//	- The difference between to levels is at least 1 and at most 3
// Otherwise they are unsafe
// Goal is to see how many reports are safe

// Part 2: The problem dampener can tolerate 1 bad level (i.e. one level can be "skipped")
// So, we need to check if:
//	1. We have already skipped a level, and
//	2. The next level follows the same rules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	None Direction = iota
	Asc
	Desc
)

func main() {
	reports := createInputMatrix()
	// reports := [][]int {
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// 	{0, 3, 2, 1},
	// 	{0, 3, 2, 5, 7},
	// 	{0, 4, 2, 3, 5},
	// }

	numReportsSafe := 0
	failed := make([][]int, 0)

	for _, report := range reports {
		if isReportSafe(report, true, None) {
			numReportsSafe++
		} else {
			failed = append(failed, report)
		}
	}

	fmt.Println("Number of total reports:", len(reports))
	fmt.Println("Number of safe reports: ", numReportsSafe)
}

func createInputMatrix() [][]int {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]int

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(text))

		for i, s := range text {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting text: ", err)
				continue
			}

			levels[i] = num
		}

		input = append(input, levels)
	}

	return input
}

/*
* Reports are safe if:
*	1. Levels are all increasing or all decreasing, AND
*	2. The difference between levels is 1 <= levels <= 3
 */
func isReportSafe(report []int, canSkip bool, direction Direction) bool {
	reportSafe := true

	// Report of length 1 should be safe (?)
	if len(report) >= 2 {
		first := report[0]
		second := report[1]

		if direction == None && first < second {
			direction = Asc
		} else if direction == None {
			direction = Desc
		}

		// If we hit a problem, we need to determine if the report can continue,
		// by removing the current OR previous level
		for i := range report {
			if i == 0 {
				continue
			}
			currIdx := i
			prevIdx := i - 1

			levelSafe := isSafe(direction, report[prevIdx], report[currIdx])

			if !levelSafe && canSkip {
				// If level is not safe, we can either skip the current or previous level
				// For the 3rd level, it is also possible to skip the first level
				skipPrevLevelSafe := prevIdx > 1 && isReportSafe(append([]int{report[prevIdx-2]}, report[currIdx:]...), false, direction)
				skipCurrLevelSafe := currIdx < len(report)-1 && isReportSafe(append([]int{report[prevIdx]}, report[currIdx+1:]...), false, direction)

				reportSafe = currIdx == len(report)-1 || skipPrevLevelSafe || skipCurrLevelSafe // If we are at the last level, we can just skip it and be done

				if currIdx == 2 && !reportSafe {
					reportSafe = isReportSafe(report[1:], false, None)
				}

				break
			} else if !levelSafe {
				reportSafe = false
				break
			}
		}
	}

	return reportSafe
}

/*
* Reports are safe if:
*	1. Levels are all increasing or all decreasing, AND
*	2. The difference between levels is 1 <= levels <= 3
 */
func isSafe(setDirection Direction, l1 int, l2 int) bool {
	var diff int
	var direction Direction

	if l1 < l2 {
		diff = l2 - l1
		direction = Asc
	} else {
		diff = l1 - l2
		direction = Desc
	}

	return direction == setDirection && diff >= 1 && diff <= 3
}

// if !levelSafe && !levelSkipped { if prevIdx >= 0 { var removePrevLevel bool
// 		if currIdx == 2 {
// 			// If we are at idx 2 and are considering removing idx 1, we may have to change direction
// 			var oneRemoveSafe bool
// 			var zeroRemoveSafe bool
// 			direction := report[2] < report[3] // Look ahead to see direction we should prioritize
//
// 			oneRemoveSafe = isSafe(direction, report[0], report[2])
// 			if !oneRemoveSafe {
// 				zeroRemoveSafe = isSafe(direction, report[1], report[2])
// 			}
//
// 			removePrevLevel = oneRemoveSafe || zeroRemoveSafe
// 			if removePrevLevel {
// 				increasing = direction
// 			}
// 		} else if currIdx > 2 {
// 			// Check if prev removed is valid
// 			removePrevLevel = isSafe(increasing, report[prevIdx-1], report[currIdx])
// 		}
//
// 		if !removePrevLevel {
// 			report[currIdx] = report[prevIdx] // If the previous level cannot be removed, we are forced to remove ourself
// 		}
// 	}
//
// 	levelSkipped = true
// } else if !levelSafe {
// 	// If report is not safe and we have already performed a skip,
// 	// break out of the loop
// 	reportSafe = false
// 	break
// }
