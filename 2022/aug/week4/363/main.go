package main
import (
	"fmt"
	"math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	sumMatrix := make([][]int, len(matrix))
	for i, _ := range sumMatrix {
		sumMatrix[i] = make([]int, len(matrix[0]))
		rowSum := int(0)
		for j, _ := range sumMatrix[i] {
			rowSum += matrix[i][j]
			sumMatrix[i][j] = rowSum

			if i > 0 {
				sumMatrix[i][j] += sumMatrix[i - 1][j]
			}
		}
	}

	result := math.MinInt32
	for y, _ := range matrix {
		for x, _ := range matrix[y] {
			helper(y, x, k, sumMatrix, &result) 
		}
	}

	return result
}

func helper(y int, x int, target int, sumMatrix [][]int, result *int) {
	if y >= len(sumMatrix) || x >= len(sumMatrix[0]) {
		return
	}

	for i := y; i < len(sumMatrix); i++ {
		for j := x; j < len(sumMatrix[0]); j++ {
			count := sumMatrix[i][j]
			if y >= 1 && x >= 1 {
				count = count - (sumMatrix[i][x - 1] + sumMatrix[y - 1][j] - sumMatrix[y - 1][x - 1])
			} else if y <= 0 && x >= 1 {
				count = count - sumMatrix[i][x - 1]
			} else if y >= 1 && x <= 0 {
				count = count - sumMatrix[y - 1][j]
			}

			if count <= target && count > *result {
				*result = count
			}
		}
	}
}

func main() {
	// result: 2
	// matrix := [][]int{{1,0,1},{0,-2,3}}
	// k := int(2)

	// result: 3
	// matrix := [][]int{{2,2,-1}}
	// k := int(3)

	// result: 15
	matrix := [][]int{
		{-2,-3,8,-9,4,7,6},
		{4,0,-8,10,4,-2,-1},
		{-2,1,0,5,6,7,1},
		{23,-8,9,-4,-5,7,6},
		{12,-12,18,-19,14,17,16},
		{-2,-3,8,-9,4,7,6},
		{-2,-3,8,-9,4,7,6},
	}
	k := int(15)

	// result: 
	// matrix := [][]int{}
	// k := int(0)

	result := maxSumSubmatrix(matrix, k)
	fmt.Printf("result = %v\n", result)
}

