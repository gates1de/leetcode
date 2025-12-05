package main

import (
	"fmt"
)

func countPartitions(nums []int) int {
	result := int(0)
	sum := int(0)
	for _, num := range nums {
		sum += num
	}

	leftSum := int(0)
	rightSum := sum
	for i := 0; i < len(nums) - 1; i++ {
		num := nums[i]
		leftSum += num
		rightSum -= num

		if (leftSum - rightSum) % 2 == 0 {
			result++
		}
	}

	return result
}

func main() {
	// result:4 
	// nums := []int{10,10,3,7,6}

	// result: 0
	// nums := []int{1,2,2}

	// result: 3
	nums := []int{2,4,6,8}

	// result: 
	// nums := []int{}

	result := countPartitions(nums)
	fmt.Printf("result = %v\n", result)
}

