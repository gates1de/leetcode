package main

import (
	"fmt"
)

func numberOfSubmatrices(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	sumX := make([]int, cols)
	sumY := make([]int, cols)
	result := int(0)
	for i := range rows {
		rx := int(0)
		ry := int(0)

		for j := range cols {
			switch grid[i][j] {
			case 'X':
				rx++
			case 'Y':
				ry++
			}

			sumX[j] += rx
			sumY[j] += ry

			if sumX[j] > 0 && sumX[j] == sumY[j] {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 3
	// grid := [][]byte{{'X','Y','.'},{'Y','.','.'}}

	// result: 0
	// grid := [][]byte{{'X','X'},{'X','Y'}}

	// result: 0
	grid := [][]byte{{'.','.'},{'.','.'}}

	// result: 
	// grid := [][]byte{}

	result := numberOfSubmatrices(grid)
	fmt.Printf("result = %v\n", result)
}

