package main

import (
	"fmt"
)

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	for range 4 {
		for i := range n / 2 {
			for j := range (n + 1) / 2 {
				mat[i][j], mat[n - 1 - j][i], mat[n - 1 - i][n - 1 - j], mat[j][n - 1 - i] = 
				mat[n - 1 - j][i], mat[n - 1 - i][n - 1 - j], mat[j][n - 1 - i], mat[i][j]
			}
		}

		if isEqual(mat, target) {
			return true
		}
	}

	return false
}

func isEqual(mat, target [][]int) bool {
	n := len(mat)
	for i := range n {
		for j := range n {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	// result: true
	// mat := [][]int{{0,1},{1,0}}
	// target := [][]int{{1,0},{0,1}}

	// result: false
	// mat := [][]int{{0,1},{1,1}}
	// target := [][]int{{1,0},{0,1}}

	// result: true
	mat := [][]int{{0,0,0},{0,1,0},{1,1,1}}
	target := [][]int{{1,1,1},{0,1,0},{0,0,0}}

	// result: 
	// mat := [][]int{}
	// target := [][]int{}

	result := findRotation(mat, target)
	fmt.Printf("result = %v\n", result)
}

