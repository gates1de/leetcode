package main

import (
	"fmt"
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	result := int(0)
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	for i := 0; i < len(points)-1; i++ {
		pointA := points[i]
		xMin := pointA[0] - 1
		xMax := math.MaxInt32
		yMin := math.MinInt32
		yMax := pointA[1] + 1

		for j := i + 1; j < len(points); j++ {
			pointB := points[j]

			if pointB[0] > xMin && pointB[0] < xMax &&
				pointB[1] > yMin && pointB[1] < yMax {
				result++
				xMin = pointB[0]
				yMin = pointB[1]
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

