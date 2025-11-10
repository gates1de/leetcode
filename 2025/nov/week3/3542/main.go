package main

import (
	"fmt"
)

func minOperations(nums []int) int {
	s := []int{}
	result := int(0)

	for _, num := range nums {
		for len(s) > 0 && s[len(s) - 1] > num {
			s = s[:len(s) - 1]
		}

		if num == 0 {
			continue
		}

		if len(s) == 0 || s[len(s) - 1] < num {
			result++
			s = append(s, num)
		}
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{0,2}

	// result: 3
	// nums := []int{3,1,2,1}

	// result: 4
	nums := []int{1,2,1,2,1,2}

	// result: 
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}

