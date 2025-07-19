package main
import (
	"fmt"
)

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	result := int(0)
	for _, num := range nums {
		num %= k

		for prev := 0; prev < k; prev++ {
			dp[prev][num] = dp[num][prev] + 1
			result = max(result, dp[prev][num])
		}
	}

	return result
}

func main() {
	// result: 5
	// nums := []int{1,2,3,4,5}
	// k := int(2)

	// result: 4
	nums := []int{1,4,2,3,1,4}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maximumLength(nums, k)
	fmt.Printf("result = %v\n", result)
}

