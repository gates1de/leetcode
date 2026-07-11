package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a][0] == intervals[b][0] {
			return intervals[a][1] > intervals[b][1]
		}
		return intervals[a][0] < intervals[b][0]
	})

	result := int(0)
	right := int(-1)
	for _, interval := range intervals {
		if interval[1] > right {
			result++
			right = interval[1]
		}
	}

	return result
}

func main() {
	// result: 2
	// intervals := [][]int{{1,4},{3,6},{2,8}}

	// result: 1
	// intervals := [][]int{{1,4},{2,3}}

	// result: 1
	// intervals := [][]int{{0,100000}}

	// result: 4
	// intervals := [][]int{{1,4},{4,5},{5,7},{8,9}}

	// result: 1
	// intervals := [][]int{{8,16},{1,2},{1,100},{2,5},{50,100},{3,19},{1,101}}

	// result: 2
	intervals := [][]int{{3,10},{4,10},{5,11}}

	// result: 
	// intervals := [][]int{}

	result := removeCoveredIntervals(intervals)
	fmt.Printf("result = %v\n", result)
}
