package main

import (
	"fmt"
)

func minSwaps(grid [][]int) int {
	zeroCount := make([]int, len(grid))
	pointer := int(0)
	for _, row := range grid {
		count := int(0)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == 0 {
				count++
			} else {
				break
			}
		}
		zeroCount[pointer] = count
		pointer++
	}

	result := int(0)
	for i := 0; i < len(zeroCount) - 1; i++ {
		if zeroCount[i] >= len(zeroCount) - i - 1 {
			continue
		}

		found := false
		for j := i + 1; j < len(zeroCount); j++ {
			if zeroCount[j] >= len(zeroCount) - i - 1 {
				found = true
				result += j - i

				for j > i{
					temp := zeroCount[j-1];
					zeroCount[j - 1] = zeroCount[j]
					zeroCount[j] = temp
					j--
				}

				break
			}
		}

		if !found {
			return -1
		}
	}

	return result
}

func main() {
	// result:
	// grid := [][]int{{0,0,1},{1,1,0},{1,0,0}}

	// result: -1
	// grid := [][]int{{0,1,1,0},{0,1,1,0},{0,1,1,0},{0,1,1,0}}

	// result: 0
	grid := [][]int{{1,0,0},{1,1,0},{1,1,1}}

	// result:
	// grid := [][]int{}

	result := minSwaps(grid)
	fmt.Printf("result = %v\n", result)
}

