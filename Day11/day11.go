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
	stones := readInput("day11_input.txt")
	// stones := readInput("day11_example.txt")
	// stones := []string{"125", "17"}
	blinks := 25

	for i := 0; i < blinks; i++ {
		toAdd := blink(stones)
		offset := 0
		for idx, str := range toAdd {
			// fmt.Println("Inserting", str, "at index", idx)
			stones = slices.Insert(stones, idx+offset, str)
			offset++
		}

		// fmt.Println(stones)
	}

	fmt.Println(len(stones))
}

func blink(stones []string) map[int]string {
	toAdd := make(map[int]string)
	for i := 0; i < len(stones); i++ {
		// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
		if stones[i] == "0" {
			stones[i] = "1"
		} else if len(stones[i])%2 == 0 {
			// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
			// The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved
			// on the right stone
			half := len(stones[i]) / 2
			first := strings.TrimLeft(stones[i][:half], "0")
			second := strings.TrimLeft(stones[i][half:], "0")

			// If we ended up trimming everything, it must have been all 0s, so insert 0
			if first == "" {
				first = "0"
			}
			if second == "" {
				second = "0"
			}

			// In order to not modify the slice as we are iterating, add the second half to a map to append after
			stones[i] = first
			toAdd[i+1] = second
		} else {
			// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is
			// engraved on the new stone.
			num, err := strconv.Atoi(stones[i])
			if err != nil {
				panic(err)
			}
			stones[i] = strconv.Itoa(num * 2024)
		}
	}
	return toAdd
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
