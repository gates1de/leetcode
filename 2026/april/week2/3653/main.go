package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func xorAfterQueries(nums []int, queries [][]int) int {
	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for i := l; i <= r; i += k {
			nums[i] = int((int64(nums[i]) * int64(v)) % modulo)
		}
	}

	result := int(0)
	for _, x := range nums {
		result ^= x
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{1,1,1}
	// queries := [][]int{{0,2,1,4}}

	// result: 31
	nums := []int{2,3,1,5,4}
	queries := [][]int{{1,4,2,3},{0,2,1,2}}

	// result: 0
	// nums := []int{}
	// queries := [][]int{}

	result := xorAfterQueries(nums, queries)
	fmt.Printf("result = %v\n", result)
}

