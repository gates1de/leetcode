package main

import (
	"fmt"
)

func findGCD(nums []int) int {
	minValue := nums[0]
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	for maxValue != 0 {
		minValue, maxValue = maxValue, minValue%maxValue
	}

	return minValue
}

func main() {
	// result: 2
	// nums := []int{2,5,6,9,10}

	// result: 1
	// nums := []int{7,5,6,8,3}

	// result: 3
	nums := []int{3,3}

	// result: 
	// nums := []int{}

	result := findGCD(nums)
	fmt.Printf("result = %v\n", result)
}
