package main

import (
	"fmt"
)

func minNumberOperations(target []int) int {
	result := target[0]
	for i := 1; i < len(target); i++ {
		result += max(target[i] - target[i - 1], 0)
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
	// result: 3
	// target := []int{1,2,3,2,1}

	// result: 4
	// target := []int{3,1,1,2}

	// result: 7
	target := []int{3,1,5,4,2}

	// result: 
	// target := []int{}

	result := minNumberOperations(target)
	fmt.Printf("result = %v\n", result)
}

