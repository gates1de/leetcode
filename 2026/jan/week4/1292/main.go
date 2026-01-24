package main

import (
	"fmt"
)

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	P := make([][]int, m + 1)
	for i := range P {
		P[i] = make([]int, n + 1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			P[i][j] = P[i - 1][j] + P[i][j - 1] - P[i - 1][j - 1] + mat[i - 1][j - 1]
		}
	}

	r := min(m, n)
	result := int(0)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for c := result + 1; c <= r; c++ {
				if i + c - 1 <= m && j + c - 1 <= n && getRect(P, i, j, i + c - 1, j + c - 1) <= threshold {
					result = c
				} else {
					break
				}
			}
		}
	}

	return result
}

func getRect(P [][]int, x1, y1, x2, y2 int) int {
	return P[x2][y2] - P[x1 - 1][y2] - P[x2][y1 - 1] + P[x1 - 1][y1 - 1]
}

func main() {
	// result: 2
	// mat := [][]int{{1,1,3,2,4,3,2},{1,1,3,2,4,3,2},{1,1,3,2,4,3,2}}
	// threshold := int(4)

	// result: 0
	mat := [][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}
	threshold := int(1)

	// result: 
	// mat := [][]int{}
	// threshold := int()

	result := maxSideLength(mat, threshold)
	fmt.Printf("result = %v\n", result)
}

