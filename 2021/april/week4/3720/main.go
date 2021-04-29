package main
import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	// total number of "matrices" aka number of spiral cycles
	m := n / 2
	for i := 0; i < m; i++ {
		limit := n - i - 1
		for j := i; j < limit; j++ {
			// copy right and move top to right
			copyRight := matrix[j][n-i-1]
			matrix[j][n-i-1] = matrix[i][j]

			// copy bottom and move copy of right to bottom
			copyBottom := matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = copyRight

			// copy left and move copy of bottom to left
			copyLeft := matrix[n-j-1][i]
			matrix[n-j-1][i] = copyBottom

			// move copy of left to top
			matrix[i][j] = copyLeft
		}
	}
}

// A little slow & Use more memory
func myAnswer(matrix [][]int) {
	n := len(matrix)
	copyMatrix := make([][]int, len(matrix))
	for i, nums := range matrix {
		copyNums := make([]int, len(nums))
		for j, num := range nums {
			copyNums[j] = num
		}
		copyMatrix[i] = copyNums
	}
	// fmt.Printf("copyMatrix = %v\n", copyMatrix)

	for i, nums := range copyMatrix {
		for j, num := range nums {
			matrix[j][n - 1 - i] = num
		}
	}
}

func main() {
	// result: [[7,4,1],[8,5,2],[9,6,3]]
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// result: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	// result: 
	// matrix := [][]int{}

	rotate(matrix)
	fmt.Printf("result = %v\n", matrix)
}

