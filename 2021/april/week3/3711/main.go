package main
import (
	"fmt"
	"math"
)

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n - 1; j++ {
			matrix[i][j + 1] += matrix[i][j]
		}
	}
	result := int(0)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			preSums := map[int]int{0: 1}
			s := 0
			for k := 0; k < m; k++ {
				if i > 0 {
					s += matrix[k][j] - matrix[k][i - 1]
				} else {
					s += matrix[k][j]
				}
				result += preSums[s - target]
				preSums[s]++
			}
		}
	}
	return result
}

// Time Limit Exceeded
func ngSolution(matrix [][]int, target int) int {
	if len(matrix) == 0 {
		return 0
	}
	result := int(0)
	maxRow := len(matrix)
	maxCol := len(matrix[0])
	for i := 0; i < maxRow; i++ { // row start
		for j := 1; j < maxRow - i + 1; j++ { // row size
			for m := 0; m < maxCol; m++ { // column start
				for n := 1; n < maxCol - m + 1; n++ { // column size
					// fmt.Printf("[%v, %v, %v, %v]\n", i, j, m, n)
					// printMatrix(i, j, m, n, matrix)
					matrixSum := sum(i, j, m, n, matrix)
					if matrixSum == target {
						result++
					}
					// fmt.Printf("sum = %v\n", matrixSum)
				}
			}
		}
	}
	return result
}

func sum(rowStart, rowSize, columnStart, columnSize int, matrix [][]int) int {
	// fmt.Printf("rowStart = %v, rowSize = %v, columnStart = %v, columnsSize = %v\n", rowStart, rowSize, columnStart, columnSize)
    if columnStart + columnSize > len(matrix[0]) {
		return math.MaxInt32
	}
    if rowStart + rowSize > len(matrix) {
		return math.MaxInt32
	}
	result := int(0)
	for i := rowStart; i < rowStart + rowSize; i++ {
		for j := columnStart; j < columnStart + columnSize; j++ {
			// fmt.Printf("%v ", matrix[i][j])
			result += matrix[i][j]
		}
		// fmt.Printf("\n")
	}
	// fmt.Printf("------------------------------------\n")
	return result
}

func main() {
	// result: 4
	matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	target := int(0)

	// result: 5
	// matrix := [][]int{{1, -1}, {-1, 1}}
	// target := int(0)

	// result: 0
	// matrix := [][]int{{904}}
	// target := int(0)

	// result: 
	// matrix := [][]int{}
	// target := int(0)

	result := numSubmatrixSumTarget(matrix, target)
	fmt.Printf("result = %v\n", result)
}

