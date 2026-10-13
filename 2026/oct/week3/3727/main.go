package main

import (
	"fmt"
	"sort"
)

func maxAlternatingSum(nums []int) int64 {
	squares := make([]int64, len(nums))
	for i, value := range nums {
		value64 := int64(value)
		squares[i] = value64 * value64
	}
	sort.Slice(squares, func(i, j int) bool {
		return squares[i] < squares[j]
	})

	negativeCount := len(nums) / 2
	result := int64(0)
	for i, square := range squares {
		if i < negativeCount {
			result -= square
		} else {
			result += square
		}
	}
	return result
}

func main() {
	// result: 12
	// nums := []int{1,2,3}

	// result: 16
	nums := []int{1, -1, 2, -2, 3, -3}

	// result: 0
	// nums := []int{}

	result := maxAlternatingSum(nums)
	fmt.Printf("result = %v\n", result)
}
