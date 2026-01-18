package main

import (
	"fmt"
)

func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	rowSum := make([][]int, m)

	for i := range rowSum {
		rowSum[i] = make([]int, n)
		rowSum[i][0] = grid[i][0]
		for j := 1; j < n; j++ {
			rowSum[i][j] = rowSum[i][j - 1] + grid[i][j]
		}
	}

	colSum := make([][]int, m)
	for i := range colSum {
		colSum[i] = make([]int, n)
	}

	for j := range n {
		colSum[0][j] = grid[0][j]
		for i := 1; i < m; i++ {
			colSum[i][j] = colSum[i - 1][j] + grid[i][j]
		}
	}

	for edge := min(m, n); edge >= 2; edge-- {
		for i := 0; i + edge <= m; i++ {
			for j := 0; j + edge <= n; j++ {
				stdSum := rowSum[i][j+edge-1]
				if j > 0 {
					stdSum -= rowSum[i][j-1]
				}

				check := true
				for ii := i + 1; ii < i + edge; ii++ {
					sum := rowSum[ii][j+edge-1]
					if j > 0 {
						sum -= rowSum[ii][j-1]
					}
					if sum != stdSum {
						check = false
						break
					}
				}

				if !check {
					continue
				}

				for jj := j; jj < j+edge; jj++ {
					sum := colSum[i+edge-1][jj]
					if i > 0 {
						sum -= colSum[i-1][jj]
					}
					if sum != stdSum {
						check = false
						break
					}
				}

				if !check {
					continue
				}

				d1, d2 := 0, 0
				for k := 0; k < edge; k++ {
					d1 += grid[i+k][j+k]
					d2 += grid[i+k][j+edge-1-k]
				}
				if d1 == stdSum && d2 == stdSum {
					return edge
				}
			}
		}
	}

	return 1
}

func main() {
	// result: 3
	// grid := [][]int{{7,1,4,5,6},{2,5,1,6,4},{1,6,4,3,1},{6,4,3,3,1}}

	// result: 2
	grid := [][]int{{5,1,3,1},{9,3,3,1},{1,3,3,8}}

	// result: 
	// grid := [][]int{}

	result := largestMagicSquare(grid)
	fmt.Printf("result = %v\n", result)
}

