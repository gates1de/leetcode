package main
import (
	"fmt"
)

type NumMatrix struct {
	SumMat [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	sumMat := make([][]int, m)
	for i, _ := range sumMat {
		sumMat[i] = make([]int, n)
		copy(sumMat[i], matrix[i])

		if i == 0 {
			for j := 1; j < n; j++ {
				sumMat[i][j] += sumMat[i][j - 1]
			}
			continue
		}

		for j, _ := range sumMat[i] {
			sumMat[i][j] += sumMat[i - 1][j]

			if j == 0 {
				continue
			}

			for _, num := range matrix[i][:j] {
				sumMat[i][j] += num
			}
		}
	}

	return NumMatrix{sumMat}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	m := len(this.SumMat)
	n := len(this.SumMat[0])
	if row1 < 0 || row1 >= m || row2 >= m || col1 < 0 || col1 >= n || col2 >= n {
		return -1
	}

	if row1 == 0 && col1 == 0 {
		return this.SumMat[row2][col2]
	}

	sum := this.SumMat[row2][col2]
	if row1 == 0 {
		return sum - this.SumMat[row2][col1 - 1]
	}
	if col1 == 0 {
		return sum - this.SumMat[row1 - 1][col2]
	}

	return sum - this.SumMat[row1 - 1][col2] - this.SumMat[row2][col1 - 1] + this.SumMat[row1 - 1][col1 - 1]
}

func main() {
	// result: [null, 8, 11, 12]
	// matrix := [][]int{
	// 	{3, 0, 1, 4, 2},
	// 	{5, 6, 3, 2, 1},
	// 	{1, 2, 0, 1, 5},
	// 	{4, 1, 0, 1, 7},
	// 	{1, 0, 3, 0, 5},
	// }
	// controls := [][]int{
	// 	{2, 1, 4, 3},
	// 	{1, 1, 2, 2},
	// 	{1, 2, 2, 4},
	// }

	// result: [null, 38, 12, 26, 5]
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	controls := [][]int{
		{0, 0, 4, 3},
		{0, 1, 2, 2},
		{1, 0, 2, 4},
		{4, 4, 4, 4},
	}

	// result: 
	// matrix := [][]int{}
	// controls := [][]int{}

	obj := Constructor(matrix)
	for _, control := range controls {
		fmt.Printf("sum = %v\n", obj.SumRegion(control[0],control[1],control[2],control[3]))
	}
}

