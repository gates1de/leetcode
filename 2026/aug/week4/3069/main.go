package main

import (
	"fmt"
)

func resultArray(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for _, num := range nums[2:] {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, num)
		} else {
			arr2 = append(arr2, num)
		}
	}

	return append(arr1, arr2...)
}

func main() {
	// result: [2,3,1]
	// nums := []int{2,1,3}

	// result: [5,3,4,8]
	nums := []int{5, 4, 3, 8}

	// result: []
	// nums := []int{}

	result := resultArray(nums)
	fmt.Printf("result = %v\n", result)
}
