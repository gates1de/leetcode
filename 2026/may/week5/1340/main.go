package main

import (
	"fmt"
)

func maxJumps(arr []int, d int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		if memo[i] != 0 {
			return memo[i]
		}

		best := int(1)

		for step := 1; step <= d && i - step >= 0; step++ {
			j := i - step
			if arr[j] >= arr[i] {
				break
			}

			best = max(best, 1 + dfs(j))
		}

		for step := 1; step <= d && i + step < n; step++ {
			j := i + step
			if arr[j] >= arr[i] {
				break
			}

			best = max(best, 1 + dfs(j))
		}

		memo[i] = best
		return best
	}

	result := int(1)
	for i := range n {
		result = max(result, dfs(i))
	}

	return result
}

func main() {
	// result: 4
	// arr := []int{6,4,14,6,8,13,9,7,10,6,12}
	// d := int(2)

	// result: 1
	// arr := []int{3,3,3,3,3}
	// d := int(3)

	// result: 7
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	d := int(1)

	// result: 0
	// arr := []int{}
	// d := int(0)

	result := maxJumps(arr, d)
	fmt.Printf("result = %v\n", result)
}
