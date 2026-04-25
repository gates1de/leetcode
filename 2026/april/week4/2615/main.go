package main

import (
	"fmt"
)

func distance(nums []int) []int64 {
	n := len(nums)
	groups := make(map[int][]int)
	for i := range n {
		groups[nums[i]] = append(groups[nums[i]], i)
	}

	result := make([]int64, n)
	for _, group := range groups {
		total := int64(0)
		for _, idx := range group {
			total += int64(idx)
		}

		prefixTotal := int64(0)
		for i, idx := range group {
			result[idx] = total - prefixTotal * 2 + int64(idx) * int64(2 * i - len(group))
			prefixTotal += int64(idx)
		}
	}

	return result
}

func main() {
	// result: [5,0,3,4,0]
	// nums := []int{1,3,1,1,2}

	// result: [0,0,0]
	nums := []int{0,5,3}

	// result: 
	// nums := []int{}

	result := distance(nums)
	fmt.Printf("result = %v\n", result)
}

