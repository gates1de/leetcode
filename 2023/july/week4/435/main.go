package main
import (
	"fmt"
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	result := int(0)
	k := math.MinInt32
	for _, interval := range intervals {
		x := interval[0]
		y := interval[1]
		if x >= k {
			k = y
		} else {
			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// intervals := [][]int{{1,2},{2,3},{3,4},{1,3}}

	// result: 2
	// intervals := [][]int{{1,2},{1,2},{1,2}}

	// result: 0
	intervals := [][]int{{1,2},{2,3}}

	// result: 
	// intervals := [][]int{}

	result := eraseOverlapIntervals(intervals)
	fmt.Printf("result = %v\n", result)
}

