package main
import (
	"fmt"
)

func maxPoints(points [][]int) int64 {
	cols := len(points[0])
	currRows := make([]int64, cols)
	prevRows := make([]int64, cols)

	for _, row := range points {
		runningMax := int64(0)

		for col := 0; col < cols; col++ {
			runningMax = max(runningMax - 1, prevRows[col])
			currRows[col] = runningMax
		}
		
		for col := cols - 1; col >= 0; col-- {
			runningMax = max(runningMax - 1, prevRows[col])
			currRows[col] = max(currRows[col], runningMax) + int64(row[col])
		}

		prevRows = currRows
	}

	result := int64(0)
	for col := 0; col < cols; col++ {
		result = max(result, prevRows[col])
	}

	return result
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 9
	// points := [][]int{{1,2,3},{1,5,1},{3,1,1}}

	// result: 11
	points := [][]int{{1,5},{2,3},{4,2}}

	// result: 
	// points := [][]int{}

	result := maxPoints(points)
	fmt.Printf("result = %v\n", result)
}

