package main

import (
	"fmt"
	"sort"
)

func minimizeSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	return min(nums[n-3]-nums[0], nums[n-2]-nums[1], nums[n-1]-nums[2])
}

func main() {
	// result: 3
	// nums := []int{1,4,7,8,5}

	// result: 0
	nums := []int{1, 4, 3}

	// result: 0
	// nums := []int{}

	result := minimizeSum(nums)
	fmt.Printf("result = %v\n", result)
}
