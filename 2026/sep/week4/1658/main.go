package main

import (
	"fmt"
)

func minOperations(nums []int, x int) int {
	total := 0
	for _, value := range nums {
		total += value
	}

	target := total - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}

	left := 0
	windowSum := 0
	longest := 0
	for right, value := range nums {
		windowSum += value
		for windowSum > target {
			windowSum -= nums[left]
			left++
		}
		if windowSum == target && right-left+1 > longest {
			longest = right - left + 1
		}
	}

	if longest == 0 {
		return -1
	}
	return len(nums) - longest
}

func main() {
	// result: 2
	// nums := []int{1,1,4,2,3}
	// x := int(5)

	// result: -1
	// nums := []int{5,6,7,8,9}
	// x := int(4)

	// result: 5
	nums := []int{3, 2, 20, 1, 1, 3}
	x := int(10)

	result := minOperations(nums, x)
	fmt.Printf("result = %v\n", result)
}
