package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = original[i*n : (i+1)*n]
	}
	return result
}

func main() {
	// result: [[1,2],[3,4]]
	// original := []int{1, 2, 3, 4}
	// m := int(2)
	// n := int(2)

	// result: [[1,2,3]]
	// original := []int{1,2,3}
	// m := int(1)
	// n := int(3)

	// result: []
	original := []int{1, 2}
	m := int(1)
	n := int(1)

	// result: []
	// original := []int{}
	// m := int(0)
	// n := int(0)

	result := construct2DArray(original, m, n)
	fmt.Printf("result = %v\n", result)
}
