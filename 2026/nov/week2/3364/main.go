package main

import (
	"fmt"
)

func minimumSumSubarray(nums []int, l int, r int) int {
	result := 0
	found := false

	for start := range nums {
		sum := 0
		for end := start; end < min(len(nums), start+r); end++ {
			sum += nums[end]
			length := end - start + 1
			if length >= l && sum > 0 && (!found || sum < result) {
				result = sum
				found = true
			}
		}
	}

	if !found {
		return -1
	}
	return result
}

func main() {
	// result: 1
	// nums := []int{3, -2, 1, 4}
	// l := 2
	// r := 3

	// result: -1
	// nums := []int{-2, 2, -3, 1}
	// l := 2
	// r := 3

	// result: 3
	nums := []int{1, 2, 3, 4}
	l := 2
	r := 4

	// result: 0
	// nums := []int{}
	// l := 0
	// r := 0

	result := minimumSumSubarray(nums, l, r)
	fmt.Printf("result = %v\n", result)
}
