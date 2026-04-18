package main

import (
	"fmt"
)

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	pos := make(map[int]int)

	for i := -n; i < n; i++ {
		if i >= 0 {
			left[i] = pos[nums[i]]
		}
		pos[nums[(i + n) % n]] = i
	}

	pos = make(map[int]int)
	for i := 2 * n - 1; i >= 0; i-- {
		if i < n {
			right[i] = pos[nums[i]]
		}
		pos[nums[i % n]] = i
	}

	for i := range len(queries) {
		x := queries[i]
		if x - left[x] == n {
			queries[i] = -1
		} else {
			queries[i] = min(x - left[x], right[x] - x)
		}
	}

	return queries
}

func main() {
	// result: [2,-1,3]
	// nums := []int{1,3,1,4,1,3,2}
	// queries := []int{0,3,5}

	// result: [-1,-1,-1,-1]
	nums := []int{1,2,3,4}
	queries := []int{0,1,2,3}

	// result: []
	// nums := []int{}
	// queries := []int{}

	result := solveQueries(nums, queries)
	fmt.Printf("result = %v\n", result)
}

