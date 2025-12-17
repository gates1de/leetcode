package main

import (
	"fmt"
)

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][3]int64, k + 1)
	for j := 1; j <= k; j++ {
		dp[j][1] = int64(-prices[0])
		dp[j][2] = int64(prices[0])
	}

	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], max(dp[j][1] + int64(prices[i]), dp[j][2] - int64(prices[i])))
			dp[j][1] = max(dp[j][1], dp[j - 1][0] - int64(prices[i]))
			dp[j][2] = max(dp[j][2], dp[j - 1][0] + int64(prices[i]))
		}
	}

	return dp[k][0]
}

func main() {
	// result: 14
	// prices := []int{1,7,9,8,2}
	// k := int(2)

	// result: 36
	prices := []int{12,16,19,19,8,1,19,13,9}
	k := int(3)

	// result: 
	// prices := []int{}
	// k := int(0)

	result := maximumProfit(prices, k)
	fmt.Printf("result = %v\n", result)
}

