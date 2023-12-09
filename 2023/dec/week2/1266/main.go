package main
import (
	"fmt"
)

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	result := int(0)
	pre := points[0]
	for _, point := range points[1:] {
		xDistance := abs(pre[0], point[0])
		yDistance := abs(pre[1], point[1])

		if xDistance > yDistance {
			result += yDistance + (xDistance - yDistance)
		} else {
			result += xDistance + (yDistance - xDistance)
		}

		pre = point
	}

	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 7
	// points := [][]int{{1,1},{3,4},{-1,0}}

	// result: 5
	points := [][]int{{3,2},{-2,2}}

	// result: 
	// points := [][]int{}

	result := minTimeToVisitAllPoints(points)
	fmt.Printf("result = %v\n", result)
}

