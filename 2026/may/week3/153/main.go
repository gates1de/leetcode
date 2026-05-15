package main

import (
	"fmt"
)

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	low := int(0)
	high := int(n - 1)
	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			high--
		}
	}

	return nums[low]
}

func main() {
	// result: 1
	// nums := []int{1, 3, 5}

	// result: 0
	// nums := []int{2, 2, 2, 0, 1}

	// result: 0
	// nums := []int{4, 5, 6, 7, 0, 1, 2}

	// result: -3300
	// nums := []int{400,500,600,700,800,900,1000,1100,2200,-3300,-440,-85,-78,-77}

	// result: 1
	// nums := []int{3, 3, 1, 3}

	// result: 1
	nums := []int{10, 1, 10, 10, 10}

	// result:
	// nums := []int{}

	result := findMin(nums)
	fmt.Printf("result = %v\n", result)
}
