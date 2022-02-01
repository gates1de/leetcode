package main
import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := int(10000)
	result := int(0)
	for _, price := range prices {
		if min > price {
			min = price
		}

		profit := price - min
		if result < profit {
			result = profit
		}
	}

    return result
}

func main() {
	// result: 5
	// prices := []int{7, 1, 5, 3, 6, 4}

	// result: 6
	// prices := []int{7, 1, 5, 3, 6, 4, 5, 7}

	// result: 0
	// prices := []int{7, 6, 4, 3, 1}

	// result: 5
	// prices := []int{7, 6, 4, 3, 8, 1}

	// result: 0
	// prices := []int{1}

	// result: 6
	// prices := []int{1, 7}

	// result: 10
	prices := []int{1,2,3,4,5,6,7,8,9,11}

	// result: 
	// prices := []int{}

	result := maxProfit(prices)
	fmt.Printf("result = %v\n", result)
}

