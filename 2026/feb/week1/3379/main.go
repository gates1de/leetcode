package main

import (
	"fmt"
)

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, num := range nums {
		result[i] = nums[((i + num) % n + n) % n]
	}
	return result
}

func main() {
	// result: [1,1,1,3]
	// nums := []int{3,-2,1,1}

	// result: [-1,-1,4]
	nums := []int{-1,4,-1}

	// result: []
	// nums := []int{}

	result := constructTransformedArray(nums)
	fmt.Printf("result = %v\n", result)
}

