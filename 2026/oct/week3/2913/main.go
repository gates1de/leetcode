package main

import (
	"fmt"
)

func sumCounts(nums []int) int {
	result := 0
	for left := range nums {
		seen := [101]bool{}
		distinct := 0
		for _, value := range nums[left:] {
			if !seen[value] {
				seen[value] = true
				distinct++
			}
			result += distinct * distinct
		}
	}
	return result
}

func main() {
	// result: 15
	// nums := []int{1,2,1}

	// result: 3
	nums := []int{1, 1}

	// result: 0
	// nums := []int{}

	result := sumCounts(nums)
	fmt.Printf("result = %v\n", result)
}
