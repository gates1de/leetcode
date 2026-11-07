package main

import (
	"fmt"
	"sort"
)

func countServers(n int, logs [][]int, x int, queries []int) []int {
	sort.Slice(logs, func(i int, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	type query struct {
		time  int
		index int
	}
	sortedQueries := make([]query, len(queries))
	for index, time := range queries {
		sortedQueries[index] = query{time: time, index: index}
	}
	sort.Slice(sortedQueries, func(i int, j int) bool {
		return sortedQueries[i].time < sortedQueries[j].time
	})

	result := make([]int, len(queries))
	counts := make([]int, n+1)
	active := 0
	left := 0
	right := 0

	for _, currentQuery := range sortedQueries {
		for right < len(logs) && logs[right][1] <= currentQuery.time {
			server := logs[right][0]
			if counts[server] == 0 {
				active++
			}
			counts[server]++
			right++
		}

		start := currentQuery.time - x
		for left < right && logs[left][1] < start {
			server := logs[left][0]
			counts[server]--
			if counts[server] == 0 {
				active--
			}
			left++
		}

		result[currentQuery.index] = n - active
	}

	return result
}

func main() {
	// result: [1, 2]
	// n := 3
	// logs := [][]int{{1, 3}, {2, 6}, {1, 5}}
	// x := 5
	// queries := []int{10, 11}

	// result: [0, 1]
	n := 3
	logs := [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}}
	x := 2
	queries := []int{3, 4}

	// result: []
	// n := 0
	// logs := [][]int{}
	// x := 0
	// queries := []int{}

	result := countServers(n, logs, x, queries)
	fmt.Printf("result = %v\n", result)
}
