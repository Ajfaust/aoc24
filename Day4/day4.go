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

func main() {
	wordSearchGrid := createInputMatrix()
	// wordSearchGrid := []string{
	// 	"MMMSXXMASM",
	// 	"MSAMXMSMSA",
	// 	"AMXSXMAAMM",
	// 	"MSAMASMSMX",
	// 	"XMASAMXAMM",
	// 	"XXAMMXXAMA",
	// 	"SMSMSASXSS",
	// 	"SAXAMASAAA",
	// 	"MAMMMXMMMM",
	// 	"MXMXAXMASX",
	// }

	sum := 0
	boundX := len(wordSearchGrid[0])
	boundY := len(wordSearchGrid)

	for i := range wordSearchGrid {
		for j, c := range wordSearchGrid[i] {
			if c != 'X' && c != 'S' {
				continue
			}

			// Check right
			if j+3 < boundX {
				word := ""
				for r := j; r < j+4; r++ {
					word += string(wordSearchGrid[i][r])
				}
				if word == "XMAS" || word == "SAMX" {
					sum++
				}

				// Check right and down diagonal
				if i+3 < boundY {
					dwnRght := ""
					for k := 0; k < 4; k++ {
						dwnRght += string(wordSearchGrid[i+k][j+k])
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
					word += string(wordSearchGrid[k][j])
				}
				if word == "XMAS" || word == "SAMX" {
					sum++
				}
				// Check down and left
				if j-3 >= 0 {
					dwnLft := ""
					for m := 0; m < 4; m++ {
						dwnLft += string(wordSearchGrid[i+m][j-m])
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
