package main

import (
	"fmt"
)

func minBitwiseArray(nums []int) []int {
	for i, num := range nums {
		newNum := int(-1)
		d := int(1)
		for (num & d) != 0 {
			newNum = num - d
			d <<= 1
		}

		nums[i] = newNum
	}

	return nums
}

func main() {
	// result: [-1,1,4,3]
	// nums := []int{2,3,5,7}

	// result: [9,12,15]
	nums := []int{11,13,31}

	// result: []
	// nums := []int{}

	result := minBitwiseArray(nums)
	fmt.Printf("result = %v\n", result)
}

