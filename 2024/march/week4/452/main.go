package main
import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func (a, b int) bool {
		return points[a][0] < points[b][0]
	})

	stack := make([][]int, 0, len(points))
	for _, point := range points {
		slen := len(stack)
		if slen > 0 && stack[slen - 1][1] >= point[0] {
			topStack := stack[slen - 1]
			stack = stack[:slen - 1]
			lastStartPoint := topStack[0]
			lastEndPoint := topStack[1]
			stack = append(stack, []int{max(point[0], lastStartPoint), min(point[1], lastEndPoint)})
		} else {
			stack = append(stack, point)
		}
	}

	return len(stack)
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
	// result: 2
	// points := [][]int{{10,16},{2,8},{1,6},{7,12}}

	// result: 4
	// points := [][]int{{1,2},{3,4},{5,6},{7,8}}

	// result: 2
	points := [][]int{{1,2},{2,3},{3,4},{4,5}}

	// result: 
	// points := [][]int{}

	result := findMinArrowShots(points)
	fmt.Printf("result = %v\n", result)
}

