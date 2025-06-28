package main

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	if n <= k {
		return nums
	}

	vals := make([][]int, n)
	for i := 0; i < n; i++ {
		vals[i] = []int{i, nums[i]}
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i][1] > vals[j][1]
	})
	sort.Slice(vals[:k], func(i, j int) bool {
		return vals[i][0] < vals[j][0]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = vals[i][1]
	}

	return result
}

func main() {
	// result: [3,3]
	// nums := []int{2,1,3,3}
	// k := int(2)

	// result: [-1,3,4]
	// nums := []int{-1,-2,3,4}
	// k := int(3)

	// result: [3,4]
	nums := []int{3, 4, 3, 3}
	k := int(2)

	// result: []
	// nums := []int{}
	// k := int(0)

	result := maxSubsequence(nums, k)
	fmt.Printf("result = %v\n", result)
}
