package main

import (
	"fmt"
)

func kLengthApart(nums []int, k int) bool {
	count := int(k)
	for _, num := range nums {
		if num == 0 {
			count++
			continue
		}
		if count < k {
			return false
		}
		count = 0
	}
	return true
}

func main() {
	// nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	// k := int(2)

	// nums := []int{1, 0, 0, 1, 0, 1}
	// k := int(2)

	// nums := []int{1, 1, 1, 1, 1}
	// k := int(0)

	nums := []int{0, 1, 0, 1}
	k := int(1)

	result := kLengthApart(nums, k)
	fmt.Printf("result = %v\n", result)
}

