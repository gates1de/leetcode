package main

import (
	"fmt"
)

func areSimilar(mat [][]int, k int) bool {
	m := len(mat)
	n := len(mat[0])
	k %= n

	for i := range m {
		for j := range n {
			if mat[i][j] != mat[i][(j + k) % n] {
				return false
			}
		}
	}

	return true
}

func main() {
	// result: false
	// mat := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	// k := int(4)

	// result: true
	// mat := [][]int{{1,2,1,2},{5,5,5,5},{6,3,6,3}}
	// k := int(2)

	// result: true
	mat := [][]int{{2,2},{2,2}}
	k := int(3)

	// result: false
	// mat := [][]int{}
	// k := int(0)

	result := areSimilar(mat, k)
	fmt.Printf("result = %v\n", result)
}

