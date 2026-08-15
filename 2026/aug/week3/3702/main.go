package main

import (
	"fmt"
)

func longestSubsequence(nums []int) int {
	xor := 0
	hasNonZero := false
	for _, num := range nums {
		xor ^= num
		if num != 0 {
			hasNonZero = true
		}
	}

	if xor != 0 {
		return len(nums)
	}
	if hasNonZero {
		return len(nums) - 1
	}

	return 0
}

func main() {
	// result: 2
	// nums := []int{1,2,3}

	// result: 3
	nums := []int{2, 3, 4}

	// result:
	// nums := []int{}

	result := longestSubsequence(nums)
	fmt.Printf("result = %v\n", result)
}
