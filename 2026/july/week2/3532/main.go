package main

import (
	"fmt"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	component := make([]int, n)
	current := int(0)
	component[0] = current

	for i := 1; i < n; i++ {
		if nums[i] - nums[i - 1] > maxDiff {
			current++
		}
		component[i] = current
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		result[i] = component[query[0]] == component[query[1]]
	}

	return result
}

func main() {
	// result: [true, false]
	// n := int(2)
	// nums := []int{1, 3}
	// maxDiff := int(1)
	// queries := [][]int{{0, 0}, {0, 1}}

	// result: [false, false, true, true]
	n := int(4)
	nums := []int{2, 5, 6, 8}
	maxDiff := int(2)
	queries := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}

	// result: []
	// maxDiff := int(0)
	// queries := [][]int{}

	result := pathExistenceQueries(n, nums, maxDiff, queries)
	fmt.Printf("result = %v\n", result)
}
