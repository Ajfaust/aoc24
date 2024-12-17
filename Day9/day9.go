// Day 9

// PART 1
// We are given an array of numbers, representing number of file blocks and free space blocks, alternating
// (e.g. 12345 represents a block of files size 1, 2 blocks of free space, file size 3 blocks, etc.).
// The goal is to create the disk this map represents, and then fill in the empty blocks by moving in file blocks starting at the end of the disk.
// Each file also has an ID associated with it, which is the block's index after the disk is created but before they are rearranged.
// After the disk is rearranged, we need to calculate the checksum, which is calculated by multiplying the file id with it's sorted position and
// then taking the sum of all the products.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// disk := readInput("day9_example.txt")
	disk := readInput("day9_input.txt")
	formatDisk(disk)
}

func formatDisk(disk []int) {
	a := 0
	b := len(disk) - 1

	// Two indexes, one at each end. A look for free space starting at beginning, b looks for file blocks starting at end
	for {
		for disk[a] >= 0 {
			a++
		}

		for disk[b] < 0 {
			b--
		}

		if a >= b {
			break
		}

		swapBlocks(disk, a, b)
	}

	// Calculate checksum
	checksum := 0
	for i, id := range disk {
		if id == -1 {
			break
		}

		checksum += i * id
	}

	fmt.Println(checksum)
}

func swapBlocks(disk []int, a int, b int) {
	temp := disk[a]
	disk[a] = disk[b]
	disk[b] = temp
}

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result []int
	nextId := 0
	index := 0

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		} else {
			num := int(c - '0')

			if index%2 == 0 {
				for range num {
					result = append(result, nextId)
				}
				nextId++
			} else {
				for range num {
					result = append(result, -1) // -1 will indicate space since we know file cannot have an ID of less than 0
				}
			}

			index++
		}
	}

	return result
}
