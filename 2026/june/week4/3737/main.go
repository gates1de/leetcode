package main

import (
	"fmt"
)

func countMajoritySubarrays(nums []int, target int) int {
	result := int(0)
	n := len(nums)

	for i := range n {
		sum := int(0)
		for j := i; j < n; j++ {
			if nums[j] == target {
				sum++
			} else {
				sum--
			}
			if sum > 0 {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 5
	// nums := []int{1,2,2,3}
	// target := int(2)

	// result: 1
	// nums := []int{1,1,1,1}
	// target := int(1)

	// result: 0
	nums := []int{1,2,3}
	target := int(4)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := countMajoritySubarrays(nums, target)
	fmt.Printf("result = %v\n", result)
}
