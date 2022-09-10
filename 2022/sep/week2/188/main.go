package main
import (
	"fmt"
	"math"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	if k > n / 2 {
		result := int(0)
		for i := 1; i < n; i++ {
			result += max(prices[i] - prices[i - 1], 0)
		}
		return result
	}

	buys := make([]int, k + 1)
	sells := make([]int, k + 1)
	for i := 0; i <= k; i++ {
		buys[i] = math.MinInt32
		sells[i] = 0
	}

	current := int(0)
	for i := 0; i < n; i++ {
		current = prices[i]
		for j := k; j > 0; j-- {
			sells[j] = max(sells[j], buys[j] + current)
			buys[j] = max(buys[j], sells[j - 1] - current)
		}
	}
	return sells[k]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// k := int(2)
	// prices := []int{2,4,1}

	// result: 7
	// k := int(3)
	// prices := []int{3,2,6,5,0,3}

	// result: 
	// k := int(0)
	// prices := []int{}

	result := maxProfit(k, prices)
	fmt.Printf("result = %v\n", result)
}

