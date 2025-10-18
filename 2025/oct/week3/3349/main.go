package main

import (
	"fmt"
)

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if n < k {
		return false
	}

	length := int(1)
	pre := int(0)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i - 1] {
			length++
			continue
		}

		if max(length / 2, min(length, pre)) >= k {
			return true
		}

		pre = length
		length = 1
	}

	return max(length / 2, min(length, pre)) >= k
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: true
	// nums := []int{2,5,7,8,9,2,3,4,3,1}
	// k := int(3)

	// result: false
	nums := []int{1,2,3,4,4,4,4,5,6,7}
	k := int(5)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := hasIncreasingSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}

