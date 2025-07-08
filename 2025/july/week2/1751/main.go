package main

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(a, b int) bool {
		return events[a][0] < events[b][0]
	})
	dp := make([][]int, k + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}

	return dfs(0, 0, -1, events, k, dp)
}

func dfs(currentIndex int, count int, prevEndingTime int, events [][]int, k int, dp [][]int) int {
	if currentIndex == len(events) || count == k {
		return 0
	}

	if prevEndingTime >= events[currentIndex][0] {
		return dfs(currentIndex + 1, count, prevEndingTime, events, k, dp)
    }

	if dp[count][currentIndex] != -1 {
		return dp[count][currentIndex]
	}

	result := max(
		dfs(currentIndex + 1, count, prevEndingTime, events, k, dp),
		dfs(currentIndex + 1, count + 1, events[currentIndex][1], events, k, dp) + events[currentIndex][2],
	)
	dp[count][currentIndex] = result
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 7
	// events := [][]int{{1,2,4},{3,4,3},{2,3,1}}
	// k := int(2)

	// result: 10
	// events := [][]int{{1,2,4},{3,4,3},{2,3,10}}
	// k := int(2)

	// result: 9
	events := [][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}
	k := int(3)

	// result: 
	// events := [][]int{}
	// k := int(0)

	result := maxValue(events, k)
	fmt.Printf("result = %v\n", result)
}

