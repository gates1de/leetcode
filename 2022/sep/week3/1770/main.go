package main
import (
	"fmt"
)

func maximumScore(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)

	if n == 0 || m == 0 {
		return 0
	}

	dp := make([]int, m + 1)

	for opeCount := m - 1; opeCount >= 0; opeCount-- {
		nextRow := make([]int, m + 1)
		copy(nextRow, dp)
		for left := opeCount; left >= 0; left-- {
			mul := multipliers[opeCount]
			dp[left] = max(mul * nums[left] + nextRow[left + 1], mul * nums[n - 1 - (opeCount - left)] + nextRow[left])
        }
	}

	return dp[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 14
	// nums := []int{1,2,3}
	// multipliers := []int{3,2,1}

	// result: 102
	nums := []int{-5,-3,-3,-2,7,1}
	multipliers := []int{-10,-5,3,4,6}

	// result: 
	// nums := []int{}
	// multipliers := []int{}

	result := maximumScore(nums, multipliers)
	fmt.Printf("result = %v\n", result)
}

