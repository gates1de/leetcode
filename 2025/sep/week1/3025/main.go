package main

import (
	"fmt"
)

func numberOfPairs(points [][]int) int {
	n := len(points)
	result := int(0)

	for i := 0; i < n; i++ {
		pointA := points[i]

		for j := 0; j < n; j++ {
			pointB := points[j]

			if i == j || !(pointA[0] <= pointB[0] && pointA[1] >= pointB[1]) {
				continue
			}

			if n == 2 {
				result++
				continue
			}

			isIllegal := false
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}

				pointTmp := points[k]
				isXContained := pointTmp[0] >= pointA[0] && pointTmp[0] <= pointB[0]
				isYContained := pointTmp[1] <= pointA[1] && pointTmp[1] >= pointB[1]
				if isXContained && isYContained {
					isIllegal = true
					break
				}
			}

			if !isIllegal {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 0
	// points := [][]int{{1,1},{2,2},{3,3}}

	// result: 2
	// points := [][]int{{6,2},{4,4},{2,6}}

	// result: 2
	points := [][]int{{3,1},{1,3},{1,1}}

	// result: 
	// points := [][]int{}

	result := numberOfPairs(points)
	fmt.Printf("result = %v\n", result)
}

