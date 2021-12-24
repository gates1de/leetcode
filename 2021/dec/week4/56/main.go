package main
import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	result := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		pre := result[len(result) - 1]
		start := interval[0]
		end := interval[1]
		preStart := pre[0]
		preEnd := pre[1]

		if preStart <= start && end <= preEnd {
			continue
		}

		if preEnd < start {
			result = append(result, interval)
		} else if start <= preStart && end <= preEnd {
			pre[0] = start
		} else if preStart <= start && preEnd <= end {
			pre[1] = end
		} else if start <= preStart && preEnd <= end {
			pre[0] = start
			pre[1] = end
		}
	}
	return result
}

func main() {
	// result: [[1,6],[8,10],[15,18]]
	// intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}

	// result: [[1,5]]
	// intervals := [][]int{{1,4},{4,5}}

	// result: [[0,4]]
	// intervals := [][]int{{1,3},{1,4},{0,4}}

	// result: [[0,10000]]
	intervals := [][]int{{0,10000}}

	// result: 
	// intervals := [][]int{}

	result := merge(intervals)
	fmt.Printf("result = %v\n", result)
}

