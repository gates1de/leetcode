package main

import (
	"fmt"
)

func maxSubarrayLength(nums []int, k int) int {
	result := int(0)
	start := int(0)
	freq := make(map[int]uint32)
	limit := uint32(k)

	for end, value := range nums {
		freq[value]++
		for freq[value] > limit {
			old := nums[start]
			count := freq[old] - 1
			if count == 0 {
				delete(freq, old)
			} else {
				freq[old] = count
			}
			start++
		}

		length := end - start + 1
		if length > result {
			result = length
		}
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{1,2,3,1,2,3,1,2}
	// k := int(2)

	// result: 2
	// nums := []int{1,2,1,2,1,2,1,2}
	// k := int(1)

	// result: 4
	nums := []int{5, 5, 5, 5, 5, 5, 5}
	k := int(4)

	// result:
	// nums := []int{}
	// k := int(0)

	result := maxSubarrayLength(nums, k)
	fmt.Printf("result = %v\n", result)
}
