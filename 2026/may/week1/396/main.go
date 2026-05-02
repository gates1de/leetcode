package main

import (
	"fmt"
)

func maxRotateFunction(nums []int) int {
	sum := int(0)
	for _, num := range nums {
		sum += num
	}

	tmp := int(0)
	for i, num := range nums {
		tmp += i * num
	}

	result := tmp
	for i := len(nums) - 1; i > 0; i-- {
		tmp += sum - len(nums) * nums[i]
		result = max(result, tmp)
	}

	return result
}

func main() {
	// result: 26
	// nums := []int{4,3,2,6}

	// result: 0
	nums := []int{100}

	// result: 
	// nums := []int{}

	result := maxRotateFunction(nums)
	fmt.Printf("result = %v\n", result)
}

