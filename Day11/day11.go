// DAY 11

// PART 1
// We are given a list of stones, all of which change by the following rules every time we blink:
// - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
//   The left half of the digits are engraved on the new left stone, and the right half of the digits are
//   engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is
//    engraved on the new stone.
// Our goal is to find how many stones we end up with after 25 blinks.
//
// We can do this by running through the list 25 times and modifying it according to the rules
// Due to the nature of the rules, it may be easier to keep the list as strings, and convert to ints
// when we need to multiply.

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// stones := readInput("day11_example.txt")
	stones := []string{"125", "17"}
	blinks := 25

	for i := 0; i < blinks; i++ {
		blink(stones)
	}

	fmt.Println(len(stones))
}

func blink(stones []string) {
	for i := range stones {
		// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
		if stones[i] == "0" {
			stones[i] = "1"
			fmt.Println(stones)
		} else if len(stones[i])%2 == 0 {
			// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
			// The left half of the digits are engraved on the new left stone, and the right half of the digits are
			half := len(stones[i]) / 2
			first := stones[i][:half]
			second := stones[i][half:]
			fmt.Println(first, second)
			stones[i] = first
			if i == len(stones)-1 {
				stones = append(stones, second)
			} else {
				stones = slices.Insert(stones, i+1, second)
			}
			i++
			fmt.Println(stones, i)
		} else {
			// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is
			// engraved on the new stone.
			num, err := strconv.Atoi(stones[i])
			if err != nil {
				panic(err)
			}
			stones[i] = strconv.Itoa(num * 2024)
			fmt.Println(stones)
		}
	}
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input = strings.Split(scanner.Text(), " ")

	return input
}
