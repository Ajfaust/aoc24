// Day 9

// PART 1
// We are given an array of numbers, representing number of file blocks and free space blocks, alternating
// (e.g. 12345 represents a block of files size 1, 2 blocks of free space, file size 3 blocks, etc.).
// The goal is to create the disk this map represents, and then fill in the empty blocks by moving in file blocks starting at the end of the disk.
// Each file also has an ID associated with it, which is the block's index after the disk is created but before they are rearranged.
// After the disk is rearranged, we need to calculate the checksum, which is calculated by multiplying the file id with it's sorted position and
// then taking the sum of all the products.

// PART 2
// Part 2 requires the same formatting, except we have to move the entire file at once into a free space
// While reading in the input file, we can also create a size map that will allow us to search the disk for
// the next available free space before that file ID

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// disk, sizeMap := readInput("day9_example.txt")
	disk, sizeMap := readInput("day9_input.txt")
	formatDisk(disk, sizeMap)
}

func formatDisk(disk []int, sizeMap map[int]int) {
	a := len(disk) - 1

	// Two indexes, one at each end. A look for free space starting at beginning, b looks for file blocks starting at end
	for a >= 0 {
		fileId := disk[a]

		for fileId < 0 {
			a--
			fileId = disk[a]
		}

		size := sizeMap[fileId]
		start, end := findFreeSpace(disk, size, fileId)
		if start > 0 {
			for i := range disk[start : end+1] {
				swapBlocks(disk, start+i, a)
				a--
			}
		} else {
			a -= sizeMap[fileId]
		}
	}

	// Calculate checksum
	checksum := 0
	for i, id := range disk {
		if id == -1 {
			continue
		}

		checksum += i * id
	}

	fmt.Println(checksum)
}

func findFreeSpace(disk []int, size int, id int) (int, int) {
	start := -1
	end := -1
	curSize := 0
	for i, num := range disk {
		if num == id {
			start = -1
			end = -1
			break
		}

		if num == -1 {
			if start < 0 {
				start = i
			}
			curSize++
		} else {
			start = -1
			end = -1
			curSize = 0
		}

		if curSize == size {
			end = i
			break
		}
	}

	return start, end
}

func swapBlocks(disk []int, a int, b int) {
	temp := disk[a]
	disk[a] = disk[b]
	disk[b] = temp
}

func readInput(filename string) ([]int, map[int]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result []int
	sizeMap := make(map[int]int)
	nextId := 0
	index := 0

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		} else {
			num := int(c - '0')

			if index%2 == 0 {
				sizeMap[nextId] = num
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

	return result, sizeMap
}
