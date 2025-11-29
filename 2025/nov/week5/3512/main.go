package main

import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	sum := int(0)
	for _, num := range nums {
		sum += num
	}
	return sum % k
}

func main() {
	// result: 4
	// nums := []int{3,9,7}
	// k := int(5)

	// result: 0
	// nums := []int{4,1,3}
	// k := int(4)

	// result: 5
	nums := []int{3,2}
	k := int(6)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

