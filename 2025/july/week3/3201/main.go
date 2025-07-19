package main
import (
	"fmt"
)

func maximumLength(nums []int) int {
	result := int(0)
	patterns := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}

	for _, pattern := range patterns {
		count := int(0)

		for _, num := range nums {
			if num%2 == pattern[count%2] {
				count++
			}
		}

		result = max(result, count)
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
	// result: 4
	// nums := []int{1,2,3,4}

	// result: 6
	// nums := []int{1,2,1,1,2,1,2}

	// result: 2
	nums := []int{1,3}

	// result: 
	// nums := []int{}

	result := maximumLength(nums)
	fmt.Printf("result = %v\n", result)
}

