package main
import (
	"fmt"
)

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	maxLen := int(0)

	for i := 2; i < n; i++ {
		start := int(0)
		end := i - 1

		for start < end {
			pairSum := arr[start] + arr[end]

			if pairSum > arr[i] {
				end--
			} else if pairSum < arr[i] {
				start++
			} else {
				dp[end][i] = dp[start][end] + 1
				maxLen = max(maxLen, dp[end][i])
				end--
				start++
			}
		}
	}

	if maxLen == 0 {
		return 0
	}

	return maxLen + 2
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// arr := []int{1,2,3,4,5,6,7,8}

	// result: 3
	arr := []int{1,3,7,11,12,14,18}

	// result: 
	// arr := []int{}

	result := lenLongestFibSubseq(arr)
	fmt.Printf("result = %v\n", result)
}

