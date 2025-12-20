package main

import (
	"fmt"
)

func getDescentPeriods(prices []int) int64 {
	result := int64(1)
	prev := int64(1)
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i - 1] - 1 {
			prev++
		} else {
			prev = 1
		}
		result += prev
	}

	return result
}

func main() {
	// result: 7
	// prices := []int{3,2,1,4}

	// result: 4
	// prices := []int{8,6,7,7}

	// result: 1
	prices := []int{1}

	// result: 
	// prices := []int{}

	result := getDescentPeriods(prices)
	fmt.Printf("result = %v\n", result)
}

