package main
import (
	"fmt"
)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	counter := make(map[string]int)

	for _, row := range matrix {
		pattern := make([]byte, len(row))

		for i, col := range row {
			if row[0] == col {
				pattern[i] = 'T'
			} else {
				pattern[i] = 'F'
			}
		}

		patternStr := string(pattern)
		counter[patternStr]++
	}

	result := int(0)
	for _, count := range counter {
		result = max(result, count)
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
	// matrix := [][]int{{0,1},{1,1}}

	// result: 2
	// matrix := [][]int{{0,1},{1,0}}

	// result: 2
	matrix := [][]int{{0,0,0},{0,0,1},{1,1,0}}

	// result: 
	// matrix := [][]int{}

	result := maxEqualRowsAfterFlips(matrix)
	fmt.Printf("result = %v\n", result)
}

