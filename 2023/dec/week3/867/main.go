package main
import (
	"fmt"
)

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return matrix
	}

	n := len(matrix[0])
	if n == 0 {
		return matrix
	}

    result := make([][]int, n)
    for i, _ := range result {
        result[i] = make([]int, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            result[i][j] = matrix[j][i]
        }
    }

    return result
}

func main() {
	// result: [[1,4,7],[2,5,8],[3,6,9]]
	// matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: [[1,4],[2,5],[3,6]]
	matrix := [][]int{{1,2,3},{4,5,6}}

	// result: 
	// matrix := [][]int{}

	result := transpose(matrix)
	fmt.Printf("result = %v\n", result)
}

