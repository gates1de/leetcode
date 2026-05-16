package main

import (
	"fmt"
)

func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, limit * 2 + 2)
	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
		a, b := nums[i], nums[j]
		if a > b {
			a, b = b, a
		}

		diff[2] += 2
		diff[a + 1]--
		diff[a + b]--
		diff[a + b + 1]++
		diff[b + limit + 1]++
	}

	result := len(nums)
	moves := int(0)
	for sum := 2; sum <= limit * 2; sum++ {
		moves += diff[sum]
		if moves < result {
			result = moves
		}
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{1,2,4,3}
	// limit := int(4)

	// result: 2
	// nums := []int{1,2,2,1}
	// limit := int(2)

	// result: 0
	nums := []int{1, 2, 1, 2}
	limit := int(2)

	// result:
	// nums := []int{}
	// limit := int(0)

	result := minMoves(nums, limit)
	fmt.Printf("result = %v\n", result)
}
