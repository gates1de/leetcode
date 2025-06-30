package main

import (
	"fmt"
)

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	result := int(0)
	for num, count := range m {
		if _, exists := m[num + 1]; exists {
			length := count + m[num + 1]
			result = max(result, length)
		}
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
	// result: 5
	// nums := []int{1, 3, 2, 2, 5, 2, 3, 7}

	// result: 2
	// nums := []int{1, 2, 3, 4}

	// result: 0
	nums := []int{1, 1, 1, 1}

	// result: 
	// nums := []int{}

	result := findLHS(nums)
	fmt.Printf("result = %v\n", result)
}

