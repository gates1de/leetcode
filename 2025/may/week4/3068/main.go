package main
import (
	"fmt"
	"math"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	dp := make([][]int64, n + 1)
	for i, _ := range dp {
		dp[i] = make([]int64, 2)
	}
	dp[n][1] = 0
	dp[n][0] = math.MinInt64

	for i := n - 1; i >= 0; i-- {
		for even := 0; even <= 1; even++ {
			performVal := dp[i + 1][even ^ 1] + int64(nums[i] ^ k)
			dontPerformVal := dp[i + 1][even] + int64(nums[i])

			dp[i][even] = max(performVal, dontPerformVal)
		}
	}

	return dp[0][1]
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// nums := []int{1,2,1}
	// k := int(3)
	// edges := [][]int{{0,1},{0,2}}

	// result: 9
	// nums := []int{2,3}
	// k := int(7)
	// edges := [][]int{{0,1}}

	// result: 42
	nums := []int{7,7,7,7,7,7}
	k := int(3)
	edges := [][]int{{0,1},{0,2},{0,3},{0,4},{0,5}}

	// result: 
	// nums := []int{}
	// k := int(0)
	// edges := [][]int{}

	result := maximumValueSum(nums, k, edges)
	fmt.Printf("result = %v\n", result)
}

