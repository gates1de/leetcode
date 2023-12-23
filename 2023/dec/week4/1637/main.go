package main
import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func (a, b int) bool {
		return points[a][0] < points[b][0]
	})
	result := int(0)

	for i := 1; i < len(points); i++ {
		result = max(result, points[i][0] - points[i - 1][0])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// points := [][]int{{8,7},{9,9},{7,4},{9,7}}

	// result: 3
	points := [][]int{{3,1},{9,0},{1,0},{1,4},{5,3},{8,8}}

	// result: 
	// points := [][]int{}

	result := maxWidthOfVerticalArea(points)
	fmt.Printf("result = %v\n", result)
}

