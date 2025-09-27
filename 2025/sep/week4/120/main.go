package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min := triangle[i + 1][j]
			if triangle[i + 1][j + 1] < min {
				min = triangle[i + 1][j + 1]
			}
			triangle[i][j] += min
		}
	}

	return triangle[0][0]
}

func main() {
	// result: 11
	// triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	// result: -10
	// triangle := [][]int{{-10}}

	// result: -1
	triangle := [][]int{{-1}, {2, 3}, {1, -1, -3}}

	// result: 
	// triangle := [][]int{}

	result := minimumTotal(triangle)
	fmt.Printf("result = %v\n", result)
}

