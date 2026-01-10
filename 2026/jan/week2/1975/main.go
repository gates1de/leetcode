package main

import (
	"fmt"
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	result := int64(0)
	minAbs := math.MaxInt32
	negativeCount := int(0)

	for _, row := range matrix {
		for _, val := range row {
			result += int64(abs(val))
			if val < 0 {
				negativeCount++
			}
			minAbs = min(minAbs, abs(val))
		}
	}

	if negativeCount % 2 != 0 {
		result -= int64(2 * minAbs)
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// matrix := [][]int{{1,-1},{-1,1}}

	// result: 16
	matrix := [][]int{{1,2,3},{-1,-2,-3},{1,2,3}}

	// result: 
	// matrix := [][]int{}

	result := maxMatrixSum(matrix)
	fmt.Printf("result = %v\n", result)
}

