package main

import (
	"fmt"
	"sort"
)

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	dp := make(map[int]int, len(nums)*2)
	result := 0

	for _, value := range nums {
		if candidate := dp[value] + 1; candidate > dp[value+1] {
			dp[value+1] = candidate
		}
		if candidate := dp[value-1] + 1; candidate > dp[value] {
			dp[value] = candidate
		}
		result = max(result, dp[value], dp[value+1])
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{2,1,5,1,1}

	// result: 1
	nums := []int{1, 4, 7, 10}

	// result: 0
	// nums := []int{}

	result := maxSelectedElements(nums)
	fmt.Printf("result = %v\n", result)
}
