package main
import (
	"fmt"
)

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	totalSum := int(0)
	for _, num := range nums {
		totalSum += num
	}

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, 2 * totalSum + 1)
	}

	dp[0][nums[0] + totalSum] = 1
	dp[0][-nums[0] + totalSum] += 1

	for i := 1; i < n; i++ {
		for sum := -totalSum; sum <= totalSum; sum++ {
			if dp[i - 1][sum + totalSum] > 0 {
				dp[i][sum + nums[i] + totalSum] += dp[i - 1][sum + totalSum]
				dp[i][sum - nums[i] + totalSum] += dp[i - 1][sum + totalSum]
			}
		}
	}

	if abs(target) > totalSum {
		return 0
	}

	return dp[n - 1][target + totalSum]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 5
	// nums := []int{1,1,1,1,1}
	// target := int(3)

	// result: 1
	nums := []int{1}
	target := int(1)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := findTargetSumWays(nums, target)
	fmt.Printf("result = %v\n", result)
}

