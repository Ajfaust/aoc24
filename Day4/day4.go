package main

import (
	"bufio"
	"fmt"
	"os"
)

// We are given a n*m grid for a word search, and need to find the amount of 'XMAS's
// We can check up, down, left, right, and diagonally, either forwards or backwards

// Right away we can go through the grid in order (left-to-right, top-to-bottom)
// And check all possible directions for both 'XMAS' and 'SAMX'
// This would mean only performing checks if we hit an X or an S
// This also means we only have to check right, down, and both diagonals

// For part 2, our search changes to find the word MAS is the shape of an X
// Same rules apply for directions

// We know what we are looking for is a 3x3 subsection of the grid, so we
// can carve out a section if we come across an M or an S and check like
// we did in part 1
// Another way is to look for the As as we know they are the middle, and then check the corners

func main() {
	wordSearchGrid := createInputMatrix()
	// wordSearchGrid := []string{
	// 	".M.S......",
	// 	"..A..MSMS.",
	// 	".M.S.MAA..",
	// 	"..A.ASMSM.",
	// 	".M.S.M....",
	// 	"..........",
	// 	"S.S.S.S.S.",
	// 	".A.A.A.A..",
	// 	"M.M.M.M.M.",
	// 	"..........",
	// }

	sum := 0
	boundX := len(wordSearchGrid[0])
	boundY := len(wordSearchGrid)

	for i := range wordSearchGrid {
		for j, c := range wordSearchGrid[i] {
			// Only check for valid answers if letter is the start or end letter
			if c != 'M' && c != 'S' {
				continue
			}

			// To find an x, we need to go down-right from the upper left, and down-left from the upper right
			// If we hit a boundary, we don't need to check
			// We only check diagonally down so we dont encounter duplicates
			if i+2 < boundY && j+2 < boundX {
				// First, if the next letter diagonal down is not an A, then fail
				if wordSearchGrid[i+1][j+1] != 'A' {
					continue
				}

				// Then we just check if the corners are S and M or M and S, respectively
				topLeft := wordSearchGrid[i][j]
				topRight := wordSearchGrid[i][j+2]
				bottomRight := wordSearchGrid[i+2][j+2]
				bottomLeft := wordSearchGrid[i+2][j]

				// fmt.Println(string(topLeft), string(bottomRight), string(topRight), string(bottomLeft))

				if (topLeft == 'M' && bottomRight == 'S' ||
					topLeft == 'S' && bottomRight == 'M') &&
					(topRight == 'M' && bottomLeft == 'S' ||
						topRight == 'S' && bottomLeft == 'M') {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func createInputMatrix() []string {
	file, err := os.Open("day4_input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func partOne(grid []string) {
	sum := 0
	boundX := len(grid[0])
	boundY := len(grid)

	for i := range grid {
		for j, c := range grid[i] {
			if c != 'X' && c != 'S' {
				continue
			}

			// Check right
			if j+3 < boundX {
				word := ""
				for r := j; r < j+4; r++ {
					word += string(grid[i][r])
				}
				if word == "XMAS" || word == "SAMX" {
					sum++
				}

				// Check right and down diagonal
				if i+3 < boundY {
					dwnRght := ""
					for k := 0; k < 4; k++ {
						dwnRght += string(grid[i+k][j+k])
					}

					if dwnRght == "XMAS" || dwnRght == "SAMX" {
						sum++
					}
				}
			}

			// Check down
			if i+3 < boundY {
				word := ""
				for k := i; k < i+4; k++ {
					word += string(grid[k][j])
				}
				if word == "XMAS" || word == "SAMX" {
					sum++
				}
				// Check down and left
				if j-3 >= 0 {
					dwnLft := ""
					for m := 0; m < 4; m++ {
						dwnLft += string(grid[i+m][j-m])
					}

					if dwnLft == "XMAS" || dwnLft == "SAMX" {
						sum++
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
