package main

// PART 1
// We are given a list of calibration results (int) along with a list of numbers
// that are missing the operators '+' and '*'. The goal is to see how many calibration
// results can be achieved by adding a combination of the missing operators back into
// the equation.
//
// An important distinction is that operators are always evaluated left-to-right, and
// not according to PEMDAS, which will make this easier

// PART 2
// We have an additional operator '||', which concatenates the two sides together. This
// makes things trickier as we may have to switch back and forth between string and int
// to concatentate.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct to hold calibration result and list of numbers
type Calibration struct {
	result  int
	numbers []int
}

func main() {
	calibrations := readInput("day7_input.txt")
	// calibrations := readInput("day7_example.txt")
	sum := 0

	for _, cal := range calibrations {
		if isEquationTrue(cal.numbers, cal.result, 1, cal.numbers[0]) {
			sum += cal.result
		}
	}

	fmt.Println(sum)
}

// Options are + and *
// Part 2 adds || (concatenate)
// Operators work left to right IN ORDER
func isEquationTrue(nums []int, target int, idx int, total int) bool {
	if total > target {
		return false
	}
	if idx == len(nums) {
		return total == target
	}

	strTotal := strconv.Itoa(total)
	strCurr := strconv.Itoa(nums[idx])
	concat, err := strconv.Atoi(strTotal + strCurr)
	if err != nil {
		panic("Error concatenating numbers")
	}

	return isEquationTrue(nums, target, idx+1, concat) || isEquationTrue(nums, target, idx+1, total+nums[idx]) || isEquationTrue(nums, target, idx+1, total*nums[idx])
}

func readInput(filename string) []Calibration {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var calibrations []Calibration

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), ":", "") // Remove the : to make parsing easier
		var nums []int

		for _, n := range strings.Split(line, " ") {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println("Unable to parse", n, "as int")
			}
			nums = append(nums, num)
		}

		equation := Calibration{
			result:  nums[0],
			numbers: nums[1:],
		}

		calibrations = append(calibrations, equation)
	}

	return calibrations
}
