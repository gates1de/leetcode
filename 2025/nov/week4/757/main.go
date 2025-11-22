package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	prev1 := intervals[0][1] - 1
	prev2 := intervals[0][1]
	result := int(2)

	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		if prev2 < start {
			prev1 = end - 1
			prev2 = end
			result += 2
		} else if prev1 < start {
			if end == prev2 {
				prev1 = end - 1
			} else {
				prev1 = end
			}

			if prev1 > prev2 {
				prev1, prev2 = prev2, prev1
			}

			result++
		}
	}

	return result
}

func main() {
	// result: 5
	// intervals := [][]int{{1,3},{3,7},{8,9}}

	// result: 3
	// intervals := [][]int{{1,3},{1,4},{2,5},{3,5}}

	// result: 5
	intervals := [][]int{{1,2},{2,3},{2,4},{4,5}}

	// result: 
	// intervals := [][]int{}

	result := intersectionSizeTwo(intervals)
	fmt.Printf("result = %v\n", result)
}

