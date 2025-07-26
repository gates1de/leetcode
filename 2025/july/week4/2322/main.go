package main

import (
	"fmt"
	"math"
)

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	e := make([][]int, n)
	for _, v := range edges {
		e[v[0]] = append(e[v[0]], v[1])
		e[v[1]] = append(e[v[1]], v[0])
	}

	sum := int(0)
	for _, x := range nums {
		sum ^= x
	}

	result := math.MaxInt32
	dfs(0, -1, sum, nums, e, &result)
	return result
}

func dfs2(x, f, oth, anc, sum int, nums []int, e [][]int, result *int) int {
	son := nums[x]
	for _, y := range e[x] {
		if y == f {
			continue
		}

		son ^= dfs2(y, x, oth, anc, sum, nums, e, result)
	}

	if f == anc {
		return son
	}

	*result = min(*result, calc(oth, son, sum ^ oth ^ son))
	return son
}

func dfs(x, f, sum int, nums []int, e [][]int, result *int) int {
	son := nums[x]
	for _, y := range e[x] {
		if y == f {
			continue
		}

		son ^= dfs(y, x, sum, nums, e, result)
	}

	for _, y := range e[x] {
		if y == f {
			dfs2(y, x, son, x, sum, nums, e, result)
		}
	}

	return son
}

func calc(part1, part2, part3 int) int {
	return max(part1, max(part2, part3)) - min(part1, min(part2, part3))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// resultult: 9
	// nums := []int{1,5,5,4,11}
	// edges := [][]int{{0,1},{1,2},{1,3},{3,4}}

	// resultult: 0
	nums := []int{5,5,2,4,4,2}
	edges := [][]int{{0,1},{1,2},{5,2},{4,3},{1,3}}

	// resultult: 
	// nums := []int{}
	// edges := [][]int{}

	resultult := minimumScore(nums, edges)
	fmt.Printf("resultult = %v\n", resultult)
}

