package main


// Input is a list of reports, which look like ints
// These reports are safe if:
//	- The levels are either all increasing or all decreasing, AND
//	- The difference between to levels is at least 1 and at most 3
// Otherwise they are unsafe
// Goal is to see how many reports are safe

import (
		"bufio"
		"fmt"
		"os"
		"strings"
		"strconv"
)

func main() {
	reports := createInputMatrix()
	
	numReportsSafe := 0

	for _, report := range reports {
		if isReportSafe(report) {
			numReportsSafe++
		}
	}
	
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
func isReportSafe(report []int) bool {
	reportSafe := true

	// Report of length 1 should be safe (?)
	if len(report) >= 2 {
		increasing := report[0] < report[1]

		for i,r := range report {
			if i == 0 {
				continue
			}

			// Check for condition 1
			if (report[i - 1] > r && increasing) || (report[i - 1] < r && !increasing) {
				reportSafe = false
				// break  // Uncomment for part 1
			}

			// Check for condition 2
			var diff int
			if report[i - 1] < r {
				diff = r - report[i - 1]
			} else {
				diff = report[i - 1] - r
			}

			if diff < 1 || diff > 3 {
				reportSafe = false
				// break // Uncomment for part 1
			}
		}
	} 

	return reportSafe
}
