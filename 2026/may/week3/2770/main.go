package main

import (
	"fmt"
	"math"
)

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if abs(nums[j] - nums[i]) <= target {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	if dp[n - 1] < 0 {
		return -1
	}
	return dp[n - 1]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 3
	// nums := []int{1, 3, 6, 4, 1, 2}
	// target := int(2)

	// result: 5
	// nums := []int{1, 3, 6, 4, 1, 2}
	// target := int(3)

	// result: -1
	nums := []int{1, 3, 6, 4, 1, 2}
	target := int(0)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := maximumJumps(nums, target)
	fmt.Printf("result = %v\n", result)
}

