package main

import (
	"fmt"
)

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	size := 2 * n + 3
	bit := make([]int64, size)

	add := func(index int, value int64) {
		for index < size {
			bit[index] += value
			index += index & -index
		}
	}

	sum := func(index int) int64 {
		result := int64(0)
		for index > 0 {
			result += bit[index]
			index -= index & -index
		}
		return result
	}

	offset := n + 1
	prefix := 0
	result := int64(0)
	add(prefix + offset, 1)

	for _, num := range nums {
		if num == target {
			prefix++
		} else {
			prefix--
		}

		result += sum(prefix + offset - 1)
		add(prefix + offset, 1)
	}

	return result
}

func main() {
	// result: 5
	// nums := []int{1,2,2,3}
	// target := int(2)

	// result: 10
	// nums := []int{1,1,1,1}
	// target := int(1)

	// result: 0
	nums := []int{1,2,3}
	target := int(4)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := countMajoritySubarrays(nums, target)
	fmt.Printf("result = %v\n", result)
}
