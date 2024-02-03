package main
import (
	"fmt"
)

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))
	for i, _ := range dp {
		dp[i] = -1
	}

	return maxSum(arr, k, dp, 0)
}

func maxSum(arr []int, k int, dp []int, start int) int {
	n := len(arr)

	if start >= n {
		return 0
	}

	if dp[start] != -1 {
		return dp[start]
	}

	currentMax := int(0)
	result := int(0)
	end := min(n, start + k)

	for i := start; i < end; i++ {
		currentMax = max(currentMax, arr[i])
		result = max(result, currentMax * (i - start + 1) + maxSum(arr, k, dp, i + 1))
	}

	dp[start] = result
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 84
	// arr := []int{1,15,7,9,2,5,10}
	// k := int(3)

	// result: 83
	// arr := []int{1,4,1,5,7,3,6,1,9,9,3}
	// k := int(4)

	// result: 1
	arr := []int{1}
	k := int(1)

	// result: 
	// arr := []int{}
	// k := int(0)

	result := maxSumAfterPartitioning(arr, k)
	fmt.Printf("result = %v\n", result)
}

