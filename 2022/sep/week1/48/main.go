package main
import (
	"fmt"
)

func rotate(matrix [][]int)  {
	n := len(matrix)
	if n <= 1 {
		return
	}

	for i := 0; i <= n / 2; i++ {
		for j := 0; j < n - 1 - (i * 2); j++ {
			up := matrix[i][i + j]
			right := matrix[i + j][n - i - 1]
			bottom := matrix[n - i - 1][n - j - i - 1]
			left := matrix[n - i - j - 1][i]

			matrix[i + j][n - i - 1] = up
			matrix[n - i - 1][n - j - i - 1] = right
			matrix[n - j - i - 1][i] = bottom
			matrix[i][i + j] = left
		}
	}
}

func main() {
	// result: [[7,4,1],[8,5,2],[9,6,3]]
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
	// matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}

	// result: 
	// matrix := [][]int{}

	rotate(matrix)
	fmt.Printf("result = %v\n", matrix)
}

