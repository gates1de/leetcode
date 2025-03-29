package main
import (
	"fmt"
	"sort"
)

func checkValidCuts(n int, rectangles [][]int) bool {
	return helper(rectangles, 0) || helper(rectangles, 1)
}

func helper(rectangles [][]int, dim int) bool {
	gapCount := int(0)

	sort.Slice(rectangles, func(a, b int) bool {
		return rectangles[a][dim] < rectangles[b][dim]
	})

	furthestEnd := rectangles[0][dim + 2]

	for i := 1; i < len(rectangles); i++ {
		rect := rectangles[i]

		if furthestEnd <= rect[dim] {
			gapCount++
		}

		furthestEnd = max(furthestEnd, rect[dim + 2])
	}

	return gapCount >= 2
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: true
	// n := int(5)
	// rectangles := [][]int{{1,0,5,2},{0,2,2,4},{3,2,5,3},{0,4,4,5}}

	// result: true
	// n := int(4)
	// rectangles := [][]int{{0,0,1,1},{2,0,3,4},{0,2,2,3},{3,0,4,3}}

	// result: false
	n := int(4)
	rectangles := [][]int{{0,2,2,4},{1,0,3,2},{2,2,3,4},{3,0,4,2},{3,2,4,4}}

	// result: 
	// n := int(0)
	// rectangles := [][]int{}

	result := checkValidCuts(n, rectangles)
	fmt.Printf("result = %v\n", result)
}

