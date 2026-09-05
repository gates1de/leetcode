package main

import (
	"fmt"
)

func minimumDeletions(nums []int) int {
	minimumIndex, maximumIndex := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[minimumIndex] {
			minimumIndex = i
		}
		if nums[i] > nums[maximumIndex] {
			maximumIndex = i
		}
	}

	if minimumIndex > maximumIndex {
		minimumIndex, maximumIndex = maximumIndex, minimumIndex
	}

	front := maximumIndex + 1
	back := len(nums) - minimumIndex
	mixed := minimumIndex + 1 + len(nums) - maximumIndex

	return min(front, min(back, mixed))
}

func main() {
	// result: 5
	// nums := []int{2,10,7,5,4,1,8,6}

	// result: 3
	// nums := []int{0,-4,19,1,8,-2,-3,5}

	// result: 1
	nums := []int{101}

	// result:
	// nums := []int{}

	result := minimumDeletions(nums)
	fmt.Printf("result = %v\n", result)
}
