package main

import (
	"fmt"
)

func missingMultiple(nums []int, k int) int {
	present := make(map[int]bool, len(nums))
	for _, num := range nums {
		present[num] = true
	}

	for multiple := k; ; multiple += k {
		if !present[multiple] {
			return multiple
		}
	}
}

func main() {
	// result: 10
	// nums := []int{8,2,3,4,6}
	// k := int(2)

	// result: 5
	nums := []int{1, 4, 7, 10, 15}
	k := int(5)

	// result:
	// nums := []int{}
	// k := int(0)

	result := missingMultiple(nums, k)
	fmt.Printf("result = %v\n", result)
}
