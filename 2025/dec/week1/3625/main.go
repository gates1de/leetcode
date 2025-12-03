package main

import (
	"fmt"
	"math"
)

func countTrapezoids(points [][]int) int {
	n := len(points)
	slopeToIntercept := make(map[float64][]float64)
	midToSlope := make(map[float64][]float64)
	result := int(0)

	for i := range n {
		x1 := points[i][0]
		y1 := points[i][1]

		for j := i + 1; j < n; j++ {
			x2 := points[j][0]
			y2 := points[j][1]
			dx := x1 - x2
			dy := y1 - y2

			var k, b float64
			if x2 == x1 {
				k = math.MaxInt32
				b = float64(x1)
			} else {
				k = float64(y2 - y1) / float64(x2 - x1)
				b = float64(y1 * dx - x1 * dy) / float64(dx)
			}

			mid := float64((x1 + x2) * 10000 + (y1 + y2))
			slopeToIntercept[k] = append(slopeToIntercept[k], b)
			midToSlope[mid] = append(midToSlope[mid], k)
		}
	}

	for _, sti := range slopeToIntercept {
		if len(sti) == 1 {
			continue
		}

		counter := make(map[float64]int)
		for _, bVal := range sti {
			counter[bVal]++
		}

		totalSum := int(0)
		for _, count := range counter {
			result += totalSum * count
			totalSum += count
		}
	}

	for _, mts := range midToSlope {
		if len(mts) == 1 {
			continue
		}

		counter := make(map[float64]int)
		for _, kVal := range mts {
			counter[kVal]++
		}

		totalSum := int(0)
		for _, count := range counter {
			result -= totalSum * count
			totalSum += count
		}
	}

	return result
}

func main() {
	// result: 2
	// points := [][]int{{-3,2},{3,0},{2,3},{3,2},{2,-3}}

	// result: 1
	points := [][]int{{0,0},{1,0},{0,1},{2,1}}

	// result: 
	// points := [][]int{}

	result := countTrapezoids(points)
	fmt.Printf("result = %v\n", result)
}

