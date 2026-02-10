package main

import (
	"fmt"
)

func longestBalanced(nums []int) int {
	result := int(0)

	for i := range len(nums) {
		odd := make(map[int]int)
		even := make(map[int]int)

		for j := i; j < len(nums); j++ {
			if nums[j] & 1 == 1 {
				odd[nums[j]]++
			} else {
				even[nums[j]]++
			}

			if len(odd) == len(even) {
				result = max(result, j - i + 1)
			}
		}
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
	// result: 4
	// nums := []int{2,5,4,3}

	// result: 5
	// nums := []int{3,2,2,5,4}

	// result: 3
	nums := []int{1,2,3,2}

	// result: 
	// nums := []int{}

	result := longestBalanced(nums)
	fmt.Printf("result = %v\n", result)
}

