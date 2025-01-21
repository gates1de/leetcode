package main
import (
	"fmt"
	"math"
)

func gridGame(grid [][]int) int64 {
	firstRowSum := int(0)
	for _, num := range grid[0] {
		firstRowSum += num
	}

	secondRowSum := int(0)
	result := int64(math.MaxInt64)
	for i := 0; i < len(grid[0]); i++ {
		firstRowSum -= grid[0][i]
		result = min(result, int64(max(firstRowSum, secondRowSum)))
		secondRowSum += grid[1][i]
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// grid := [][]int{{2,5,4},{1,5,1}}

	// result: 4
	// grid := [][]int{{3,3,1},{8,5,2}}

	// result: 7
	grid := [][]int{{1,3,1,15},{1,3,3,1}}

	// result: 
	// grid := [][]int{}

	result := gridGame(grid)
	fmt.Printf("result = %v\n", result)
}

