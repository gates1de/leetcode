package main

import (
	"fmt"
)

func missingInteger(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		sum += nums[i]
	}

	present := make(map[int]bool, len(nums))
	for _, num := range nums {
		present[num] = true
	}
	for present[sum] {
		sum++
	}

	return sum
}

func main() {
	// result: 6
	// nums := []int{1,2,3,2,5}

	// result: 15
	nums := []int{3, 4, 5, 1, 12, 14, 13}

	// result:
	// nums := []int{}

	result := missingInteger(nums)
	fmt.Printf("result = %v\n", result)
}
