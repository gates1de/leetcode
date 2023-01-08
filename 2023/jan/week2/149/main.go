package main
import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}

	result := int(2)
	for i := 0; i < n; i++ {
		counter := make(map[float64]int, n * n)
		for j := 0; j < n; j++ {
			if i != j {
				count := math.Atan2(
					float64(points[j][1] - points[i][1]),
					float64(points[j][0] - points[i][0]),
				)
				counter[count]++
			}
		}

		for _, count := range counter {
			result = max(result, count + 1)
		}
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
	// result: 3
	// points := [][]int{{1,1},{2,2},{3,3}}

	// result: 4
	points := [][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}

	// result: 
	// points := [][]int{}

	result := maxPoints(points)
	fmt.Printf("result = %v\n", result)
}

