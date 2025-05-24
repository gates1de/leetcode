package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	firstRow := false
	firstCol := false
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstCol = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	// result: [[1,0,1],[0,0,0],[1,0,1]]
	// matrix := [][]int{
	// 	{1, 1, 1},
	// 	{1, 0, 1},
	// 	{1, 1, 1},
	// }

	// result: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
	// matrix := [][]int{
	// 	{0, 1, 2, 0},
	// 	{3, 4, 5, 2},
	// 	{1, 3, 1, 5},
	// }

	// result: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
		{2, 3, 0, 5},
		{0, 1, 1, 0},
	}

	// result:
	// matrix := [][]int{
	// 	{},
	// 	{},
	// 	{},
	// }

	setZeroes(matrix)
	fmt.Printf("result = %v\n", matrix)
}
