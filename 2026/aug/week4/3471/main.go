package main

import (
	"fmt"
)

func largestInteger(nums []int, k int) int {
	count := [51]int{}
	for start := 0; start+k <= len(nums); start++ {
		seen := [51]bool{}
		for _, value := range nums[start : start+k] {
			seen[value] = true
		}
		for value, exists := range seen {
			if exists {
				count[value]++
			}
		}
	}

	for value := len(count) - 1; value >= 0; value-- {
		if count[value] == 1 {
			return value
		}
	}

	return -1
}

func main() {
	// result: 7
	// nums := []int{3,9,2,1,7}
	// k := int(3)

	// result: 3
	// nums := []int{3,9,7,2,1,7}
	// k := int(4)

	// result: -1
	nums := []int{0, 0}
	k := int(1)

	// result:
	// nums := []int{}
	// k := int(0)

	result := largestInteger(nums, k)
	fmt.Printf("result = %v\n", result)
}
