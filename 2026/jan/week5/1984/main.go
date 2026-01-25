package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	for i, num := range nums[:len(nums) - k + 1] {
		result = min(result, nums[i + k - 1] - num)
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// nums := []int{90}
	// k := int(1)

	// result: 2
	nums := []int{9,4,1,7}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minimumDifference(nums, k)
	fmt.Printf("result = %v\n", result)
}

