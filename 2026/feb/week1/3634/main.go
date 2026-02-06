package main

import (
	"fmt"
	"sort"
)

func minRemoval(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	result := n
	right := int(0)

	for left := range n {
		for right < n && int64(nums[right]) <= int64(nums[left]) * int64(k) {
			right++
		}

		result = min(result, n - (right - left))
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// nums := []int{2,1,5}
	// k := int(2)

	// result: 2
	// nums := []int{1,6,2,9}
	// k := int(3)

	// result: 0
	nums := []int{4,6}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minRemoval(nums, k)
	fmt.Printf("result = %v\n", result)
}

