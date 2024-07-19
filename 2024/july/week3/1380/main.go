package main
import (
	"fmt"
	"math"
)

func luckyNumbers(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	rMinMax := math.MinInt32
	for i := 0; i < n; i++ {
		rMin := math.MaxInt32
		for j := 0; j < m; j++ {
			rMin = min(rMin, matrix[i][j])
		}
		rMinMax = max(rMinMax, rMin)
	}

	cMaxMin := math.MaxInt32
	for i := 0; i < m; i++ {
		cMax := math.MinInt32
		for j := 0; j < n; j++ {
			cMax = max(cMax, matrix[j][i])
		}
		cMaxMin = min(cMaxMin, cMax)
	}

	if rMinMax == cMaxMin {
		return []int{rMinMax}
	}
		
	return make([]int, 0)
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
	// result: [15]
	// matrix := [][]int{{3,7,8},{9,11,13},{15,16,17}}

	// result: [12]
	// matrix := [][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}}

	// result: [7]
	matrix := [][]int{{7,8},{1,2}}

	// result: []
	// matrix := [][]int{}

	result := luckyNumbers(matrix)
	fmt.Printf("result = %v\n", result)
}

