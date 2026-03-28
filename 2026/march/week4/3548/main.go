package main

import (
	"fmt"
	"math"
)

func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	sum := int(0)
	bottom := make([]int, 100001)
	top := make([]int, 100001)
	left := make([]int, 100001)
	right := make([]int, 100001)

	for _, rows := range grid {
		for _, col := range rows {
			sum += col
			bottom[col]++
			right[col]++
		}
	}

	sumTop := int(0)
	for i := range m - 1 {
		for j := range n {
			val := grid[i][j]
			sumTop += val
			top[val]++
			bottom[val]--
		}

		sumBottom := sum - sumTop
		if sumTop == sumBottom {
			return true
		}

		diff := math.Abs(float64(sumTop) - float64(sumBottom))
		if diff <= 100000 {
			if sumTop > sumBottom {
				if check(top, grid, 0, i, 0, n - 1, int(diff)) {
					return true
				}
			} else {
				if check(bottom, grid, i + 1, m - 1, 0, n - 1, int(diff)) {
					return true
				}
			}
		}
	}

	sumLeft := int(0)
	for j := range n - 1 {
		for i := range m {
			val := grid[i][j]
			sumLeft += val
			left[val]++
			right[val]--
		}

		sumRight := sum - sumLeft
		if sumLeft == sumRight {
			return true
		}

		diff := math.Abs(float64(sumLeft) - float64(sumRight))
		if diff <= 100000 {
			if sumLeft > sumRight {
				if check(left, grid, 0, m - 1, 0, j, int(diff)) {
					return true
				}
			} else {
				if check(right, grid, 0, m - 1, j + 1, n - 1, int(diff)) {
					return true
				}
			}
		}
	}

	return false
}

func check(freq []int, grid [][]int, r1, r2, c1, c2 int, diff int) bool {
	rows := r2 - r1 + 1
	cols := c2 - c1 + 1

	if rows * cols == 1 {
		return false
	}

	if rows == 1 {
		return grid[r1][c1] == diff || grid[r1][c2] == diff
	}
	if cols == 1 {
		return grid[r1][c1] == diff || grid[r2][c1] == diff
	}

	return freq[diff] > 0
}

func main() {
	// result: true
	// grid := [][]int{{1,4},{2,3}}

	// result: true
	// grid := [][]int{{1,2},{3,4}}

	// result: false
	grid := [][]int{{1,2,4},{2,3,5}}

	// result: false
	// grid := [][]int{}

	result := canPartitionGrid(grid)
	fmt.Printf("result = %v\n", result)
}

