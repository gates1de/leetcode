package main

import (
	"fmt"
)

func minimumDistance(nums []int) int {
	n := len(nums)
	result := n + 1

	for i := range n - 2 {
		for j := i + 1; j < n-1; j++ {
			if nums[i] != nums[j] {
				continue
			}

			for k := j + 1; k < n; k++ {
				if nums[j] == nums[k] {
					if dist := k - i; dist < result {
						result = dist
					}

					break
				}
			}
		}
	}

	if result == n + 1 {
		return -1
	}

	return result * 2
}

func main() {
	// result: 6
	// nums := []int{1,2,1,1,3}

	// result: 8
	// nums := []int{1,1,2,3,2,1,2}

	// result: -1
	nums := []int{1}

	// result: 
	// nums := []int{}

	result := minimumDistance(nums)
	fmt.Printf("result = %v\n", result)
}

