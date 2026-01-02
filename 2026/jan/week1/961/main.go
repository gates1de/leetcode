package main

import (
	"fmt"
)

func repeatedNTimes(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return num
		}
		m[num] = true
	}

	return -1
}

func main() {
	// result: 3
	// nums := []int{1,2,3,3}

	// result: 2
	// nums := []int{2,1,2,5,3,2}

	// result: 5
	nums := []int{5,1,5,2,5,3,5,4}

	// result: 
	// nums := []int{}

	result := repeatedNTimes(nums)
	fmt.Printf("result = %v\n", result)
}

