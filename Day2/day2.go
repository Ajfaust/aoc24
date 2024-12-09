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
	// reports := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// 	{0, 3, 2, 1},
	// 	{0, 3, 2, 5, 7},
	// 	{0, 4, 2, 3, 5},
	// 	{1, 2, 3, 4, 2},
	// }

	numReportsSafe := 0

	for _, report := range reports {
		if isReportSafe(report, true, None) {
			numReportsSafe++
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

		var i int
		// Iterate until we find a rule break
		for i = 0; i < len(report)-1 && reportSafe; i++ {
			reportSafe = isSafe(direction, report[i], report[i+1])
		}

		if i == len(report)-1 && canSkip {
			reportSafe = true
		} else if !reportSafe && canSkip && i < len(report)-1 { // If we are allowed to skip, brute force by skipping every level
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			// fmt.Println("Report:", report, "New Report:", reportCopy[1:])
			reportSafe = isReportSafe(reportCopy[1:], false, None)
			for j := 1; j < len(report)-1 && !reportSafe; j++ {
				if j == 1 {
					direction = None
				}
				copy(reportCopy, report)
				newReport := append(reportCopy[:j], reportCopy[j+1:]...)
				// fmt.Println("Report:", report, "New Report:", newReport)
				reportSafe = isReportSafe(newReport, false, direction)
			}

			if !reportSafe {
				copy(reportCopy, report)
				reportSafe = isReportSafe(reportCopy[:len(report)-1], false, direction)
			}

			// fmt.Println("Result:", reportSafe)
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

	// fmt.Println("l1:", l1, "l2:", l2, "setDirection:", setDirection, "direction:", direction)

	return direction == setDirection && diff >= 1 && diff <= 3
}
