package main
import (
	"fmt"
)

func diagonalSum(mat [][]int) int {
	n := len(mat)
	result := int(0)

	left := int(0)
	right := n - 1
	for i := 0; i < n; i++ {
		result += mat[i][left]
		result += mat[i][right]
		if left == right {
			result -= mat[i][left]
		}

		left++
		right--
	}

	return result
}

func main() {
	// result: 25
	// mat := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: 8
	// mat := [][]int{
	// 	{1,1,1,1},
	// 	{1,1,1,1},
	// 	{1,1,1,1},
	// 	{1,1,1,1},
	// }

	// result: 5
	mat := [][]int{{5}}

	// result: 
	// mat := [][]int{}

	result := diagonalSum(mat)
	fmt.Printf("result = %v\n", result)
}

