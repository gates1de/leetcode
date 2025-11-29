package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := int64(0)
	sum := make([]int64, k)
	result := int64(math.MinInt64)

	for i := 0; i < k; i++ {
		sum[i] = math.MaxInt64 / 2
	}
	sum[k - 1] = 0

	for i := range n {
		prefixSum += int64(nums[i])
		if prefixSum - sum[i % k] > result {
			result = prefixSum - sum[i % k]
		}
		if prefixSum < sum[i % k] {
			sum[i % k] = prefixSum
		}
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{1,2}
	// k := int(1)

	// result: -10
	// nums := []int{-1,-2,-3,-4,-5}
	// k := int(4)

	// result: 4
	nums := []int{-5,1,2,-3,4}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxSubarraySum(nums, k)
	fmt.Printf("result = %v\n", result)
}

