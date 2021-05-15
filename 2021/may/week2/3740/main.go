package main
import (
	"fmt"
)

type NumMatrix struct {
	Matrix [][]int
	Dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := NumMatrix{Matrix: matrix}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return m
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]) + 1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j + 1] = dp[i][j] + matrix[i][j]
		}
	}
	m.Dp = dp
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := int(0)
	for row := row1; row <= row2; row++ {
		result += this.Dp[row][col2 + 1] - this.Dp[row][col1]
	}
	return result
}

// Slow & Use more memory
// type NumMatrix struct {
// 	Matrix [][]int
// }
// 
// func Constructor(matrix [][]int) NumMatrix {
// 	return NumMatrix{Matrix: matrix}
// }
// 
// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
// 	result := int(0)
// 	for i := row1; i <= row2; i++ {
// 		for j := col1; j <= col2; j++ {
// 			result += this.Matrix[i][j]
// 		}
// 	}
// 	return result
// }

func main() {
	// result: [null, 8, 11, 12]
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	fmt.Printf("param1 = %v\n", obj.SumRegion(2, 1, 4, 3))
	fmt.Printf("param2 = %v\n", obj.SumRegion(1, 1, 2, 2))
	fmt.Printf("param3 = %v\n", obj.SumRegion(1, 2, 2, 4))

	// result: 
	// matrix := [][]int{}
	// row1 := int(0)
	// row2 := int(0)
	// col1 := int(0)
	// col2 := int(0)
	// obj := Constructor(matrix)
	// param_1 := obj.SumRegion(row1,col1,row2,col2)

	// result := solution()
	// fmt.Printf("result = %v\n", result)
}

