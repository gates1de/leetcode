package main

import (
	"fmt"
)

func countHillValley(nums []int) int {
	n := len(nums)
	result := int(0)

	for i := 1; i < n - 1; i++ {
		if nums[i] == nums[i - 1] {
			continue
		}

		left := int(0)
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				left = 1
				break
			} else if nums[j] < nums[i] {
				left = -1
				break
			}
		}

		right := int(0)
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				right = 1
				break
			} else if nums[j] < nums[i] {
				right = -1
				break
			}
		}

		if left == right && left != 0 {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{2,4,1,1,6,5}

	// result: 0
	nums := []int{6,6,5,5,4,1}

	// result: 
	// nums := []int{}

	result := countHillValley(nums)
	fmt.Printf("result = %v\n", result)
}

