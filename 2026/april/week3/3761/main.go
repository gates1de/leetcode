package main

import (
	"fmt"
)

func minMirrorPairDistance(nums []int) int {
	prev := make(map[int]int)
	n := len(nums)
	result := n + 1

	for i, x := range nums {
		if idx, exists := prev[x]; exists {
			if i-idx < result {
				result = i - idx
			}
		}

		prev[reverse(x)] = i
	}

	if result == n + 1 {
		return -1
	}

	return result
}

func reverse(num int) int {
	result := int(0)
	for num > 0 {
		result = result * 10 + num % 10
		num /= 10
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{12,21,45,33,54}

	// result: 1
	// nums := []int{120,21}

	// result: -1
	nums := []int{21,120}

	// result: 
	// nums := []int{}

	result := minMirrorPairDistance(nums)
	fmt.Printf("result = %v\n", result)
}

