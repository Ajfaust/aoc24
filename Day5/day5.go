package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 5 involves two inputs: a set if rules in the format X|Y, and a list of numbers representing updates
// The goal is to get all lists that have both X and Y, and X comes before Y
// Then we find the middle number and add it to a sum

func main() {
	rules, updates := parseInput()
	// rules := [][]int{
	// 	{47, 53},
	// 	{97, 13},
	// 	{97, 61},
	// 	{97, 47},
	// 	{75, 29},
	// 	{61, 13},
	// 	{75, 53},
	// 	{29, 13},
	// 	{97, 29},
	// 	{53, 29},
	// 	{61, 53},
	// 	{97, 53},
	// 	{61, 29},
	// 	{47, 13},
	// 	{75, 47},
	// 	{97, 75},
	// 	{47, 61},
	// 	{75, 61},
	// 	{47, 29},
	// 	{75, 13},
	// 	{53, 13},
	// }
	// updates := [][]int{
	// 	{75, 47, 61, 53, 29},
	// 	{97, 61, 53, 29, 13},
	// 	{75, 29, 13},
	// 	{75, 97, 47, 61, 53},
	// 	{61, 13, 29},
	// 	{97, 13, 75, 29, 47},
	// }
	sum := 0

	for _, update := range updates {
		// We will make a map of pages and their corresponding indices
		m := make(map[int]int)
		for j, page := range update {
			m[page] = j
		}

		releventRules := 0
		rulesMet := 0

		// Then we will go through the rules to see which apply
		for _, r := range rules {
			first := r[0]
			second := r[1]

			firstIdx, firstExists := m[first]
			secondIdx, secondExists := m[second]

			if !firstExists || !secondExists {
				continue
			}

			releventRules++
			if firstIdx < secondIdx {
				rulesMet++
			}
		}

		// Find the middle value if update follows all relevent rules
		if rulesMet == releventRules {
			// fmt.Println(update)
			mid := len(update) / 2
			// fmt.Println(update[mid])
			sum += update[mid]
		}
	}

	fmt.Println(sum)
}

func parseInput() ([][]int, [][]int) {
	file, err := os.Open("day5_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rules [][]int
	var updates [][]int

	splitChar := "|"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			splitChar = ","
			continue
		}

		split := strings.Split(line, splitChar)

		var slice []int
		for _, char := range split {
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			slice = append(slice, num)
		}

		if splitChar == "|" {
			rules = append(rules, slice)
		} else if splitChar == "," {
			updates = append(updates, slice)
		}
	}

	return rules, updates
}
