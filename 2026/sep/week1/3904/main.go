package main

import (
	"fmt"
)

func firstStableIndex(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	suffixMin := make([]int, len(nums))
	suffixMin[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	prefixMax := nums[0]
	for i, num := range nums {
		prefixMax = max(prefixMax, num)
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}

	return -1
}

func main() {
	// result: 3
	// nums := []int{5,0,1,4}
	// k := int(3)

	// result: -1
	// nums := []int{3,2,1}
	// k := int(1)

	// result: 0
	nums := []int{0}
	k := int(0)

	// result: 0
	// nums := []int{}
	// k := int(0)

	result := firstStableIndex(nums, k)
	fmt.Printf("result = %v\n", result)
}
