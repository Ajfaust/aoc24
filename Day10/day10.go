// DAY 10

// PART 1
// We are given a topigraphical map with numbers ranging from 0 to 9
// Our goal is to see how many ways we can get from 0 to 9, only increasing by 1

// PART 2
// We are now asked to find the number of distinct trails we can make from each 0.
// We can do so by modifying our visted map to include number of trails found
// at each location, and avoid travelling down trails we have gone down before.
// Then we just add up the number of trails we find in each direction.

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

	// Initialize a 2D array with -1. We will use this to both track visited paths and
	// their score
	visited := make([][]int, len(m))
	for i := range m {
		visited[i] = make([]int, len(m[0]))
		for j := range len(m[0]) {
			visited[i][j] = -1
		}
	}

	for i := range m {
		for j, num := range m[i] {
			if num == 0 {
				sum += getNumTrails(m, visited, j, i, 0)
			}
		}
	}

	fmt.Println(sum)
}

func getNumTrails(m [][]int, v [][]int, startX int, startY int, target int) int {
	if startX < 0 || startX >= len(m[0]) || startY < 0 || startY >= len(m) {
		return 0
	}

	num := m[startY][startX]
	if num != target {
		return 0
	}

	if num == 9 {
		return 1
	}

	// If we have analyzed this path before, no need to go through again
	if v[startY][startX] >= 0 {
		return v[startY][startX]
	}

	sum := getNumTrails(m, v, startX, startY+1, num+1) + getNumTrails(m, v, startX, startY-1, num+1) + getNumTrails(m, v, startX+1, startY, num+1) + getNumTrails(m, v, startX-1, startY, num+1)

	v[startY][startX] = sum

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
