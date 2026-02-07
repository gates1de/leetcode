package main

import (
	"fmt"
)

func minimumCost(nums []int) int {
	a := int(51)
	b := int(51)

	for i := 1; i < len(nums); i++ {
		if nums[i] < a {
			b = a
			a = nums[i]
		} else if nums[i] < b {
			b = nums[i]
		}
	}

	return nums[0] + a + b
}

func main() {
	// result: 6
	// nums := []int{1,2,3,12}

	// result: 12
	// nums := []int{5,4,3}

	// result: 12
	nums := []int{10,3,1,1}

	// result: 
	// nums := []int{}

	result := minimumCost(nums)
	fmt.Printf("result = %v\n", result)
}

