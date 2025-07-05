package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	freq := make(map[int]int)
	result := int(-1)

	for _, num := range arr {
		freq[num]++ 
	}

	for num, count := range freq {
		if num != count {
			continue
		}
		result = max(result, num)
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
	// result: 2
	// arr := []int{2,2,3,4}

	// result: 3
	// arr := []int{1,2,2,3,3,3}

	// result: -1
	arr := []int{2,2,2,3,3}

	// result: 
	// arr := []int{}

	result := findLucky(arr)
	fmt.Printf("result = %v\n", result)
}

