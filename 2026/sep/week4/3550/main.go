package main

import (
	"fmt"
)

func smallestIndex(nums []int) int {
	for index, value := range nums {
		sum := 0
		for value > 0 {
			sum += value % 10
			value /= 10
		}

		if sum == index {
			return index
		}
	}

	return -1
}

func main() {
	// result: 2
	// nums := []int{1,3,2}

	// result: 1
	// nums := []int{1,10,11}

	// result: -1
	nums := []int{1, 2, 3}

	// result: 0
	// nums := []int{}

	result := smallestIndex(nums)
	fmt.Printf("result = %v\n", result)
}
