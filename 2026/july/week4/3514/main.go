package main

import (
	"fmt"
)

const limit = 1 << 11

func uniqueXorTriplets(nums []int) int {
	present := make([]bool, limit)
	for _, value := range nums {
		present[value] = true
	}

	pairs := make([]bool, limit)
	for a := range limit {
		if !present[a] {
			continue
		}

		for b := range limit {
			if present[b] {
				pairs[a^b] = true
			}
		}
	}

	triplets := make([]bool, limit)
	for pair := range limit {
		if !pairs[pair] {
			continue
		}

		for value := range limit {
			if present[value] {
				triplets[pair^value] = true
			}
		}
	}

	result := int(0)
	for _, exists := range triplets {
		if exists {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{1,3}

	// result: 4
	nums := []int{6, 7, 8, 9}

	// result:
	// nums := []int{}

	result := uniqueXorTriplets(nums)
	fmt.Printf("result = %v\n", result)
}
