package main

import (
	"fmt"
)

func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}

	rows, cols := len(grid), len(grid[0])
	total := rows * cols
	shift := k % total
	if shift == 0 {
		return grid
	}

	for start := range gcd(total, shift) {
		current := start
		value := grid[current/cols][current%cols]

		for {
			next := (current + shift) % total
			nextValue := grid[next / cols][next % cols]
			grid[next / cols][next % cols] = value
			value = nextValue
			current = next

			if current == start {
				break
			}
		}
	}

	return grid
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	// result: [[9,1,2],[3,4,5],[6,7,8]]
	// grid := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	// k := int(1)

	// result: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
	// grid := [][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}
	// k := int(4)

	// result: [[1,2,3],[4,5,6],[7,8,9]]
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := int(9)

	// result:
	// grid := [][]int{}
	// k := int(0)

	result := shiftGrid(grid, k)
	fmt.Printf("result = %v\n", result)
}
