package main

import (
	"fmt"
)

func isGood(nums []int) bool {
	n := len(nums) - 1
	if n < 1 {
		return false
	}

	counts := make([]int, n+1)
	for _, num := range nums {
		if num < 1 || num > n {
			return false
		}
		counts[num]++
	}

	for num := 1; num < n; num++ {
		if counts[num] != 1 {
			return false
		}
	}

	return counts[n] == 2
}

func main() {
	// result: false
	// nums := []int{2, 1, 3}

	// result: true
	// nums := []int{1, 3, 3, 2}

	// result: true
	// nums := []int{1, 1}

	// result: false
	nums := []int{3, 4, 4, 1, 2, 1}

	// result:
	// nums := []int{}

	result := isGood(nums)
	fmt.Printf("result = %v\n", result)
}
