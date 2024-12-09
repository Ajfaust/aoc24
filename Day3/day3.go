package main

// Input is lines of code with a "mul()" function that takes 2 integers with 1 - 3 digits
// Input is corrupted, though
// Goal is to scan input and run any "mul()" functions we find that match the rules aboce
// and then sum up all the results

// First, we should scan for a "mul(" to find the start of a function
// Afterwards, we can scan until a ")" and read in the input
// Then, we can check that the arguments follow the rules by splitting on ','

// Part 2 adds in do() and dont() instructions, that enables and disables mul() functions, respectively
// Mul() functions are enabled at beginning

// We could add 'do()' and 'don't()' to the regex matching, and check for them as we iterate through'

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// memory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	memory := readInputFile()

	functions := getMulFunctions(memory)
	// fmt.Println(functions)
	fmt.Println(len(functions))
	sum := 0
	mulEnabled := true
	for _, f := range functions {
		if f == "do()" {
			mulEnabled = true
			continue
		} else if f == "don't()" {
			mulEnabled = false
			continue
		}

		if mulEnabled {
			nums := strings.Split(f[4:len(f)-1], ",")
			// fmt.Println(nums)
			a, errA := strconv.Atoi(nums[0])
			b, errB := strconv.Atoi(nums[1])

			if errA != nil || errB != nil {
				fmt.Println("Error converting string to numbers")
				continue
			}
			sum += a * b
		}
	}

	fmt.Println(sum)
}

func readInputFile() string {
	file, err := ioutil.ReadFile("day3_input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	return string(file)
}

func getMulFunctions(memory string) []string {
	re := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)|do\(\)|don\'t\(\)`)
	functions := re.FindAllString(memory, -1)
	return functions
}
