package main

import (
	"fmt"
)

func maxDistance(colors []int) int {
	n := len(colors)
	result := int(0)

	for i := range n {
		for j := i + 1; j < n; j++ {
			if colors[i] != colors[j] {
				result = max(result, j - i)
			}
		}
	}

	return result
}

func main() {
	// result: 3
	// colors := []int{1,1,1,6,1,1,1}

	// result: 4
	// colors := []int{1,8,3,8,3}

	// result: 1
	colors := []int{0,1}

	// result: 
	// colors := []int{}

	result := maxDistance(colors)
	fmt.Printf("result = %v\n", result)
}

