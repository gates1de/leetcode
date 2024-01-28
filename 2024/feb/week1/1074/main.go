package main
import (
	"fmt"
)

func numSubmatrixSumTarget(matrix [][]int, target int) int {
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

    result := int(0)
    for y, _ := range matrix {
        for x, _ := range matrix[y] {
            helper(y, x, target, sumMatrix, &result)
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

            if count == target {
                *result++
            }
        }
    }
}

func main() {
	// result: 4
	// matrix := [][]int{{0,1,0},{1,1,1},{0,1,0}}
	// target := int(0)

	// result: 5
	// matrix := [][]int{{1,-1},{-1,1}}
	// target := int(0)

	// result: 0
	matrix := [][]int{{904}}
	target := int(0)

	// result: 
	// matrix := [][]int{}
	// target := int(0)

	result := numSubmatrixSumTarget(matrix, target)
	fmt.Printf("result = %v\n", result)
}

