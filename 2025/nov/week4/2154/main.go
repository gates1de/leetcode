package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for _, num := range nums {
		if original == num {
			original *= 2
		}
	}

	return original
}

func main() {
	// result: 24
	// nums := []int{5,3,6,1,12}
	// original := int(3)

	// result: 4
	nums := []int{2,7,9}
	original := int(4)

	// result: 
	// nums := []int{}
	// original := int(0)

	result := findFinalValue(nums, original)
	fmt.Printf("result = %v\n", result)
}

