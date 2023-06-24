package main
import (
	"fmt"
)

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

    result := int(0)
	hold := -prices[0]
	for i := 1; i < n; i++ {
		tmp := hold
		hold = max(hold, result - prices[i])
		result = max(result, tmp + prices[i] - fee)
	}

    return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// prices := []int{1,3,2,8,4,9}
	// fee := int(2)

	// result: 6
	prices := []int{1,3,7,5,10,3}
	fee := int(3)

	// result: 
	// prices := []int{}
	// fee := int(0)

	result := maxProfit(prices, fee)
	fmt.Printf("result = %v\n", result)
}

