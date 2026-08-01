package main

import (
	"fmt"
)

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	copy(dp, nums)

	for diff := 1; diff < n; diff++ {
		for left := range n - diff {
			right := left + diff
			dp[left] = max(nums[left] - dp[left + 1], nums[right] - dp[left])
		}
	}

	if dp[0] < 0 {
		return false
	}

	return true
}

func main() {
	// result: false
	// nums := []int{1,5,2}

	// result: true
	nums := []int{1,5,233,7}

	// result: 
	// nums := []int{}

	result := predictTheWinner(nums)
	fmt.Printf("result = %v\n", result)
}
