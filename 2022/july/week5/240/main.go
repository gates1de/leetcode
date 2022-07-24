package main
import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for i, _ := range matrix {
		first := matrix[i][0]
		last := matrix[i][n - 1]
		if target < first {
			return false
		}
		if last < target {
			continue
		}

		for _, num := range matrix[i] {
			if num == target {
				return true
			}
			if num > target {
				break
			}
		}
	}

	return false
}

func main() {
	// result: true
	// matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	// target := int(5)

	// result: false
	// matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	// target := int(20)

	// result: false
	matrix := [][]int{{2, 8, 14, 22, 30},{4, 10, 16, 24, 38},{6, 12, 18, 32, 44},{20, 26, 28, 34, 48},{36, 42, 46, 52, 60}}
    target := int(21)

	// result: 
	// matrix := [][]int{}
	// target := int(0)

	result := searchMatrix(matrix, target)
	fmt.Printf("result = %v\n", result)
}

