package main

import (
	"fmt"
)

func minElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := digitSum(nums[0])
	for _, num := range nums[1:] {
		if sum := digitSum(num); sum < result {
			result = sum
		}
	}

	return result
}

func digitSum(n int) int {
	sum := int(0)
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	// result: 1
	// nums := []int{10,12,13,14}

	// result: 1
	// nums := []int{1,2,3,4}

	// result: 10
	nums := []int{999,19,199}

	// result: 
	// nums := []int{}

	result := minElement(nums)
	fmt.Printf("result = %v\n", result)
}
