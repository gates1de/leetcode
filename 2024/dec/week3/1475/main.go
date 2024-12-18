package main
import (
	"fmt"
)

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	copy(result, prices)
	stack := make([]int, 0, n)

	for i, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack) - 1]] >= price {
			result[stack[len(stack) - 1]] -= price
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	// result: [4,2,4,2,3]
	// prices := []int{8,4,6,2,3}

	// result: [1,2,3,4,5]
	// prices := []int{1,2,3,4,5}

	// result: [9,0,1,6]
	prices := []int{10,1,1,6}

	// result: []
	// prices := []int{}

	result := finalPrices(prices)
	fmt.Printf("result = %v\n", result)
}

