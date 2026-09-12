package main

import (
	"fmt"
	"sort"
)

func maximumWeight(intervals [][]int) []int {
	type interval struct {
		left   int
		right  int
		weight int64
		index  int
	}
	type choice struct {
		score int64
		ids   [4]int
		size  int
	}

	items := make([]interval, len(intervals))
	for i, current := range intervals {
		items[i] = interval{
			left:   current[0],
			right:  current[1],
			weight: int64(current[2]),
			index:  i,
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].right != items[j].right {
			return items[i].right < items[j].right
		}
		return items[i].left < items[j].left
	})

	better := func(a, b choice) bool {
		if a.score != b.score {
			return a.score > b.score
		}

		limit := min(b.size, a.size)
		for i := range limit {
			if a.ids[i] != b.ids[i] {
				return a.ids[i] < b.ids[i]
			}
		}

		return a.size < b.size
	}

	insert := func(ids [4]int, size, index int) [4]int {
		position := size
		for i := range size {
			if index < ids[i] {
				position = i
				break
			}
		}

		for i := size; i > position; i-- {
			ids[i] = ids[i-1]
		}
		ids[position] = index

		return ids
	}

	dp := make([][5]choice, len(items)+1)
	for i := range dp {
		for count := range dp[i] {
			dp[i][count].score = -1
		}
	}
	dp[0][0] = choice{}

	ends := make([]int, len(items))
	for i := range items {
		ends[i] = items[i].right
	}

	for i := 1; i <= len(items); i++ {
		for count := 0; count <= 4; count++ {
			dp[i][count] = dp[i-1][count]
		}

		current := items[i-1]
		previous := sort.Search(i-1, func(j int) bool {
			return ends[j] >= current.left
		})

		for count := 1; count <= 4; count++ {
			base := dp[previous][count-1]
			if base.score < 0 {
				continue
			}

			candidate := choice{
				score: base.score + current.weight,
				size:  count,
				ids:   insert(base.ids, count-1, current.index),
			}

			if better(candidate, dp[i][count]) {
				dp[i][count] = candidate
			}
		}
	}

	result := dp[len(items)][0]
	for count := 1; count <= 4; count++ {
		if better(dp[len(items)][count], result) {
			result = dp[len(items)][count]
		}
	}

	return result.ids[:result.size]
}

func main() {
	// result: [2,3]
	// intervals := [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}, {6, 9, 3}, {6, 7, 1}, {8, 9, 1}}

	// result: [1,3,5,6]
	intervals := [][]int{{5, 8, 1}, {6, 7, 7}, {4, 7, 3}, {9, 10, 6}, {7, 8, 2}, {11, 14, 3}, {3, 5, 5}}

	// result: []
	// intervals := [][]int{}

	result := maximumWeight(intervals)
	fmt.Printf("result = %v\n", result)
}
