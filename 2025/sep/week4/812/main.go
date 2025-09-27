package main

import (
	"fmt"
)

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	result := float64(0)

	for i := 0; i < n - 2; i++ {
		for j := i + 1; j < n - 1; j++ {
			for k := j + 1; k < n; k++ {
				x1 := float64(points[i][0])
				x2 := float64(points[j][0])
				y1 := float64(points[i][1])
				y2 := float64(points[j][1])
				x3 := float64(points[k][0])
				y3 := float64(points[k][1])

				area := 0.5 * abs(x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2))
				result = max(result, area)
			}
		}
	}

	return result
}

func abs(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2.00000
	// points := [][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}

	// result: 0.50000
	points := [][]int{{1,0},{0,0},{0,1}}

	// result: 
	// points := [][]int{}

	result := largestTriangleArea(points)
	fmt.Printf("result = %v\n", result)
}

