// DAY 10

// PART 1
// We are given a topigraphical map with numbers ranging from 0 to 9
// Our goal is to see how many ways we can get from 0 to 9, only increasing by 1

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// m := readInput("day10_example.txt")
	m := readInput("day10_input.txt")
	sum := 0

	for i := range m {
		for j, num := range m[i] {
			if num == 0 {
				// Initialize a 2D array with -1. We will use this to track visited paths
				visited := make([][]bool, len(m))
				for i := range m {
					visited[i] = make([]bool, len(m[0]))
				}
				sum += getNumTrails(m, visited, j, i, 0)
			}
		}
	}

	fmt.Println(sum)
}

func getNumTrails(m [][]int, v [][]bool, startX int, startY int, target int) int {
	if startX < 0 || startX >= len(m[0]) || startY < 0 || startY >= len(m) {
		return 0
	}

	num := m[startY][startX]
	// Don't add anything if we have already gone down this trail
	if v[startY][startX] || num != target {
		return 0
	}

	v[startY][startX] = true

	if num == 9 {
		return 1
	}

	// Trail score is the sum of all reachable 9s in each direction
	sum := getNumTrails(m, v, startX, startY+1, num+1) + getNumTrails(m, v, startX, startY-1, num+1) + getNumTrails(m, v, startX+1, startY, num+1) + getNumTrails(m, v, startX-1, startY, num+1)

	return sum
}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var m [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}

		m = append(m, row)
	}

	return m
}
