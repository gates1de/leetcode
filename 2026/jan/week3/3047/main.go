package main

import (
	"fmt"
)

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxSide := int64(0)

	for i := range n {
		for j := i + 1; j < n; j++ {
			w := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			h := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])

			if w > 0 && h > 0 {
				maxSide = max(maxSide, int64(min(w, h)))
			}
		}
	}

	return maxSide * maxSide
}

func main() {
	// result: 1
	// bottomLeft := [][]int{{1,1},{2,2},{3,1}}
	// topRight := [][]int{{3,3},{4,4},{6,6}}

	// result: 4
	// bottomLeft := [][]int{{1,1},{1,3},{1,5}}
	// topRight := [][]int{{5,5},{5,7},{5,9}}

	// result: 1
	// bottomLeft := [][]int{{1,1},{2,2},{1,2}}
	// topRight := [][]int{{3,3},{4,4},{3,4}}

	// result: 0
	bottomLeft := [][]int{{1,1},{3,3},{3,1}}
	topRight := [][]int{{2,2},{4,4},{4,2}}

	// result: 
	// bottomLeft := [][]int{}
	// topRight := [][]int{}

	result := largestSquareArea(bottomLeft, topRight)
	fmt.Printf("result = %v\n", result)
}

