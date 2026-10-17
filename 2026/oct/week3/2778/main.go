package main

import (
	"fmt"
)

func sumOfSquares(nums []int) int {
	n := len(nums)
	result := 0
	for i, value := range nums {
		index := i + 1
		if n%index == 0 {
			result += value * value
		}
	}
	return result
}

func main() {
	// result: 21
	// nums := []int{1,2,3,4}

	// result: 63
	nums := []int{2, 7, 1, 19, 18, 3}

	// result: 0
	// nums := []int{}

	result := sumOfSquares(nums)
	fmt.Printf("result = %v\n", result)
}
