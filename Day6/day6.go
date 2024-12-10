package main

import (
	"bufio"
	"fmt"
	"os"
)

// Day 6 requires us to check how many distinct positions a guard will visit on a given map
// before leaving the map
// We are given that:
//     - A # represents an obsticle
//     - A ^ represents a guard and which way they are facing
// The guards also follow a strict path, with the following requirements:
//     - If there is something directly in front of you, turn 90 degrees to the right
//     - Otherwise, take a step forward

// We are given the map as an input. We should also have a 2d slice of the same size to
// track positions visited, given that we need the number of distinct positions
// We will also need a variable to keep track of the guards direction and position
// Exit condition is when the next step goes outside the bounds of the map

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func main() {
	m, posX, posY := createInputMap("day6_input.txt")
	// m, posX, posY := createInputMap("day6_example.txt")

	direction := Up

	visited := make([][]bool, len(m))
	for i := range m {
		visited[i] = make([]bool, len(m[i]))
	}

	// fmt.Println(posY, posX)

	numPositions := 0

	// Loop until next position is out of bounds
	for posX >= 0 && posY >= 0 && posX < len(m[0]) && posY < len(m) {
		if !visited[posY][posX] {
			visited[posY][posX] = true
			numPositions++
		}
		posX, posY, direction = moveGuard(m, posX, posY, direction)
	}

	fmt.Println(numPositions)
}

// Let's read in the map, which will be of type [][]rune
// We can also use this to find the start position
func createInputMap(fileName string) ([][]rune, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input [][]rune
	startX := 0
	startY := 0
	lineNum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var row []rune
		for i, r := range line {
			row = append(row, r)
			if r == '^' {
				startX = i
				startY = lineNum
			}
		}

		input = append(input, row)
		lineNum++
	}

	return input, startX, startY
}

// The guards also follow a strict path, with the following requirements:
//   - If there is something directly in front of you, turn 90 degrees to the right
//   - Otherwise, take a step forward
func moveGuard(m [][]rune, posX int, posY int, direction Direction) (int, int, Direction) {
	switch direction {
	case Up:
		if posY-1 >= 0 && m[posY-1][posX] == '#' {
			direction = Right
			// fmt.Println("Changed direction: Right")
		} else {
			posY--
		}
		break
	case Down:
		if posY+1 < len(m) && m[posY+1][posX] == '#' {
			direction = Left
			// fmt.Println("Changed direction: Left")
		} else {
			posY++
		}
		break
	case Left:
		if posX-1 >= 0 && m[posY][posX-1] == '#' {
			direction = Up
			// fmt.Println("Changed direction: Up")
		} else {
			posX--
		}
		break
	case Right:
		if posX+1 < len(m[posY]) && m[posY][posX+1] == '#' {
			direction = Down
			// fmt.Println("Changed direction: Down")
		} else {
			posX++
		}
		break
	}

	// fmt.Println(posY, posX)
	return posX, posY, direction
}
