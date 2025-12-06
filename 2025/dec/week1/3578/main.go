package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func countPartitions(nums []int, k int) int {
	n := len(nums)
	dp := make([]int64, n + 1)
	prefix := make([]int64, n + 1)
	minQ := make([]int, 0)
	maxQ := make([]int, 0)

	dp[0] = 1
	prefix[0] = 1

	for i, j := 0, 0; i < n; i++ {
		for len(maxQ) > 0 && nums[maxQ[len(maxQ) - 1]] <= nums[i] {
			maxQ = maxQ[:len(maxQ) - 1]
		}
		maxQ = append(maxQ, i)

		for len(minQ) > 0 && nums[minQ[len(minQ) - 1]] >= nums[i] {
			minQ = minQ[:len(minQ) - 1]
		}
		minQ = append(minQ, i)

		for len(maxQ) > 0 && len(minQ) > 0 &&
			nums[maxQ[0]] - nums[minQ[0]] > k {
			if maxQ[0] == j {
				maxQ = maxQ[1:]
			}
			if minQ[0] == j {
				minQ = minQ[1:]
			}

			j++
		}

		if j > 0 {
			dp[i + 1] = (prefix[i] - prefix[j - 1] + modulo) % modulo
		} else {
			dp[i + 1] = prefix[i] % modulo
		}
		prefix[i + 1] = (prefix[i] + dp[i + 1]) % modulo
	}

	return int(dp[n])
}

func main() {
	// result: 6
	// nums := []int{9,4,1,3,7}
	// k := int(4)

	// result: 2
	nums := []int{3,3,4}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := countPartitions(nums, k)
	fmt.Printf("result = %v\n", result)
}

