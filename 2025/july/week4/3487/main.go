package main

import (
	"fmt"
	"math"
)

func maxSum(nums []int) int {
	m := make(map[int]bool)
	isEmpty := true
	minusMax := math.MinInt32
	result := int(0)

	for _, num := range nums {
		if num < 0 {
			minusMax = max(minusMax, num)
			continue
		}
		if m[num] {
			continue
		}

		m[num] = true
		isEmpty = false
		result += num
	}

	if isEmpty {
		return minusMax
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 15
	// nums := []int{1,2,3,4,5}

	// result: 1
	// nums := []int{1,1,0,1,1}

	// result: 3
	nums := []int{1,2,-1,-2,1,0,-1}

	// result: 
	// nums := []int{}

	result := maxSum(nums)
	fmt.Printf("result = %v\n", result)
}

