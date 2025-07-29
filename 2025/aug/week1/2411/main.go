package main

import (
	"fmt"
)

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	position := make([]int, 31)
	for i := range position {
		position[i] = -1
	}

	result := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		j := i

		for bit := 0; bit < 31; bit++ {
			if (nums[i] & (1 << bit)) == 0 {
				if position[bit] != -1 {
					j = max(j, position[bit])
				}
			} else {
				position[bit] = i
			}
		}

		result[i] = j - i + 1
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [3,3,2,2,1]
	// nums := []int{1,0,2,1,3}

	// result: [2,1]
	nums := []int{1,2}

	// result: []
	// nums := []int{}

	result := smallestSubarrays(nums)
	fmt.Printf("result = %v\n", result)
}

