package main
import (
	"fmt"
)

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n := len(rowSum)
	m := len(colSum)

	result := make([][]int, n)
	for i, _ := range result {
		result[i] = make([]int, m)
	}

	i := int(0)
	j := int(0)

	for i < n && j < m {
		result[i][j] = min(rowSum[i], colSum[j])

		rowSum[i] -= result[i][j]
		colSum[j] -= result[i][j]

		if rowSum[i] == 0 {
			i++
		} else {
			j++
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [[3,0],[1,7]]
	// rowSum := []int{3,8}
	// colSum := []int{4,7}

	// result: [[0,5,0],[6,1,0],[2,0,8]]
	rowSum := []int{5,7,10}
	colSum := []int{8,6,8}

	// result: []
	// rowSum := []int{}
	// colSum := []int{}

	result := restoreMatrix(rowSum, colSum)
	fmt.Printf("result = %v\n", result)
}

