package main

import (
	"fmt"
)

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	profitSum := make([]int64, n + 1)
	priceSum := make([]int64, n + 1)
	for i := range n {
		profitSum[i + 1] = profitSum[i] + int64(prices[i]) * int64(strategy[i])
		priceSum[i + 1] = priceSum[i] + int64(prices[i])
	}

	result := profitSum[n]
	for i := k - 1; i < n; i++ {
		leftProfit := profitSum[i - k + 1]
		rightProfit := profitSum[n] - profitSum[i + 1]
		changeProfit := priceSum[i + 1] - priceSum[i - k / 2 + 1]
		result = max(result, leftProfit + changeProfit + rightProfit)
	}

	return result
}

func main() {
	// result: 10
	// prices := []int{4,2,8}
	// strategy := []int{-1,0,1}
	// k := int(2)

	// result: 9
	prices := []int{5,4,3}
	strategy := []int{1,1,0}
	k := int(2)

	// result: 
	// prices := []int{}
	// strategy := []int{}
	// k := int(0)

	result := maxProfit(prices, strategy, k)
	fmt.Printf("result = %v\n", result)
}

