package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Value int
	Index int
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	pairs := make([]Pair, n)
	for i := range n {
		pairs[i] = Pair{Value: nums[i], Index: i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	pos := make([]int, n)
	values := make([]int, n)
	for i := range n {
		values[i] = pairs[i].Value
		pos[pairs[i].Index] = i
	}

	next := make([]int, n)
	right := 0
	for left := range n {
		if right < left {
			right = left
		}

		for right + 1 < n && values[right + 1] - values[left] <= maxDiff {
			right++
		}

		next[left] = right
	}

	log := int(1)
	for (1 << log) <= n {
		log++
	}

	jump := make([][]int, log)
	jump[0] = next
	for p := 1; p < log; p++ {
		jump[p] = make([]int, n)
		for i := range n {
			jump[p][i] = jump[p - 1][jump[p - 1][i]]
		}
	}

	result := make([]int, len(queries))
	for qi, query := range queries {
		u := pos[query[0]]
		v := pos[query[1]]
		if u > v {
			u, v = v, u
		}
		if u == v {
			result[qi] = 0
			continue
		}

		cur := u
		steps := 0
		for p := log - 1; p >= 0; p-- {
			nxt := jump[p][cur]
			if nxt < v {
				cur = nxt
				steps += 1 << p
			}
		}

		if jump[0][cur] < v {
			result[qi] = -1
			continue
		}

		result[qi] = steps + 1
	}

	return result
}

func main() {
	// result: [1, 1]
	// n := int(5)
	// nums := []int{1, 8, 3, 4, 2}
	// maxDiff := int(3)
	// queries := [][]int{{0, 3}, {2, 4}}

	// result: [1, 2, -1, 1]
	// n := int(5)
	// nums := []int{5, 3, 1, 9, 10}
	// maxDiff := int(2)
	// queries := [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}}

	// result: [0, -1, -1]
	n := int(3)
	nums := []int{3, 6, 1}
	maxDiff := int(1)
	queries := [][]int{{0, 0}, {0, 1}, {1, 2}}

	// result: []
	// n := int(0)
	// nums := []int{}
	// maxDiff := int(0)
	// queries := [][]int{}

	result := pathExistenceQueries(n, nums, maxDiff, queries)
	fmt.Printf("result = %v\n", result)
}
