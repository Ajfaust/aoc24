// Day 8

// We are provided a map containing lower and uppercase letters that represent antennae.
// Our goal is to count the number of antinodes we can create from these antennae that fit
// within our map.
// An antinode is described as a location that is equidistant from 2 of the same antennae
// (in both x and y). Antinodes are allowed to overlap antennae.

// We could initite a hashmap that contains the different antennae and their coordinates.
// Then we would go through each to determine corresponding antinode locations, and if they
// are in the map add 1 to a running sum.

// PART 2
// Part two includes the fact that antinode locations go as far as possible within the grid,
// following the same rules

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	m, bX, bY := readInput("day8_input.txt")
	// m, bX, bY := readInput("day8_example.txt")

	// The problem asks for UNIQUE antinode locations, so
	// we need a map containing locations that have antinodes
	// so we dont count a location twice
	antinodeMap := make([][]bool, bY)
	for i := range antinodeMap {
		antinodeMap[i] = make([]bool, bX)
	}
	numAntinodes := 0

	// Go through hashmap to find antennae coords
	for _, value := range m {
		// For each antenna, go through coords and find differences
		for i, c := range value {
			if i+1 > len(value)-1 {
				break
			}

			pairs := value[i+1:]
			for _, p := range pairs {
				coords := getAntinodeCoords(c, p, bX, bY)

				for _, c := range coords {
					if !antinodeMap[c.y][c.x] {
						numAntinodes++
						antinodeMap[c.y][c.x] = true
					}
				}
				// // If antinode coords are in bounds and another antinode is not
				// // already there, add 1 to numAntinodes
				// if isCoordInBounds(aA, bX, bY) && !antinodeMap[aA.y][aA.x] {
				// 	numAntinodes++
				// 	antinodeMap[aA.y][aA.x] = true
				// }
				// if isCoordInBounds(aB, bX, bY) && !antinodeMap[aB.y][aB.x] {
				// 	numAntinodes++
				// 	antinodeMap[aB.y][aB.x] = true
				// }
			}
		}
	}

	fmt.Println(numAntinodes)
	// for i := range antinodeMap {
	// 	fmt.Println(antinodeMap[i])
	// }
}

// Update function for part to to get all coordinates in grid range
func getAntinodeCoords(a Coordinate, b Coordinate, bX int, bY int) []Coordinate {
	diffX := a.x - b.x
	diffY := a.y - b.y

	var coords []Coordinate

	// Go up
	for isCoordInBounds(a, bX, bY) {
		coords = append(coords, a)
		a.x += diffX
		a.y += diffY
	}

	// Go down
	for isCoordInBounds(b, bX, bY) {
		coords = append(coords, b)
		b.x -= diffX
		b.y -= diffY
	}

	return coords
}

func isCoordInBounds(c Coordinate, bX int, bY int) bool {
	return c.x >= 0 && c.x < bX && c.y >= 0 && c.y < bY
}

// Read input and return map of runes and their coordinates
func readInput(filename string) (map[rune][]Coordinate, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[rune][]Coordinate)
	lineNum := 0
	boundX := 0

	for scanner.Scan() {
		line := scanner.Text()
		if boundX == 0 {
			boundX = len(line)
		}
		for i, r := range line {
			if unicode.IsDigit(r) || unicode.IsLetter(r) {
				m[r] = append(m[r], Coordinate{x: i, y: lineNum})
			}
		}

		lineNum++
	}

	boundY := lineNum

	return m, boundX, boundY
}
