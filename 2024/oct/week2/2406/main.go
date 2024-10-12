package main
import (
	"fmt"
	"math"
)

func minGroups(intervals [][]int) int {
	rangeStart := math.MaxInt32
	rangeEnd := math.MinInt32

	for _, interval := range intervals {
		rangeStart = min(rangeStart, interval[0])
		rangeEnd = max(rangeEnd, interval[1])
	}

	pointToCount := make([]int, rangeEnd + 2)

	for _, interval := range intervals {
		pointToCount[interval[0]]++
		pointToCount[interval[1] + 1]--
	}

	concurrentIntervals := int(0)
	result := int(0)

	for i := rangeStart; i <= rangeEnd; i++ {
		concurrentIntervals += pointToCount[i]
		result = max(result, concurrentIntervals)
	}

	return result
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
	// result: 3
	// intervals := [][]int{{5,10},{6,8},{1,5},{2,3},{1,10}}

	// result: 1
	intervals := [][]int{{1,3},{5,6},{8,10},{11,13}}

	// result: 
	// intervals := [][]int{}

	result := minGroups(intervals)
	fmt.Printf("result = %v\n", result)
}

