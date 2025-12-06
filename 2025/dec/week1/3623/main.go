package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countTrapezoids(points [][]int) int {
	pointMap := make(map[int]int)
	result := int(0)
	sum := int(0)

	for _, point := range points {
		y := point[1]
		pointMap[y]++
	}

	for _, point := range pointMap {
		edge := point * (point - 1) / 2
		result = (result + edge * sum) % modulo
		sum = (sum + edge) % modulo
	}

	return result
}

func main() {
	// result: 3
	// points := [][]int{{1,0},{2,0},{3,0},{2,2},{3,2}}

	// result: 1
	points := [][]int{{0,0},{1,0},{0,1},{2,1}}

	// result: 
	// points := [][]int{{0,0},{0,1},{1,1},{1,0}}

	result := countTrapezoids(points)
	fmt.Printf("result = %v\n", result)
}

