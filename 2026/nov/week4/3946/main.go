package main

import (
	"fmt"
)

func maximumSaleItems(items [][]int, budget int) int {
	if len(items) == 0 || budget == 0 {
		return 0
	}

	const maxFactor = 1500
	frequency := make([]int, maxFactor+1)
	minimumPrice := budget + 1
	for _, item := range items {
		frequency[item[0]]++
		minimumPrice = min(minimumPrice, item[1])
	}

	benefits := make([]int, maxFactor+1)
	for factor := 1; factor <= maxFactor; factor++ {
		for multiple := factor; multiple <= maxFactor; multiple += factor {
			benefits[factor] += frequency[multiple]
		}
	}

	const unavailable = -1
	dp := make([]int, budget+1)
	for cost := range dp {
		dp[cost] = unavailable
	}
	dp[0] = 0
	for _, item := range items {
		price := item[1]
		benefit := benefits[item[0]]
		for cost := budget; cost >= price; cost-- {
			if dp[cost-price] != unavailable {
				dp[cost] = max(dp[cost], dp[cost-price]+benefit)
			}
		}
	}

	result := 0
	for cost, benefit := range dp {
		if benefit != unavailable {
			result = max(result, benefit+(budget-cost)/minimumPrice)
		}
	}

	return result
}

func main() {
	// result: 4
	// items := [][]int{{6, 2},{2,6},{3,4}}
	// budget := 9

	// result: 10
	// items := [][]int{{2,4},{3,2},{4,1},{6,4},{12,4}}
	// budget := 8

	// result: 0
	items := [][]int{}
	budget := 0

	// result: 0
	// items := [][]int{}
	// budget := 0

	result := maximumSaleItems(items, budget)
	fmt.Printf("result = %v\n", result)
}
