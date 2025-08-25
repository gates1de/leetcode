package main

import (
	"fmt"
)

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	result := make([]int, m * n)

	index := int(0)
	for i := 0; i < n + m - 1; i++ {
		tmp := make([]int, 0)

		row := i - n + 1
		col := n - 1
		if i < n {
			row = 0
			col = i
		}

		for row < m && col > -1 {
			tmp = append(tmp, mat[row][col])
			row++
			col--
		}

		if i % 2 == 0 {
			reverse(tmp)
		}

		for j := 0; j < len(tmp); j++ {
			result[index] = tmp[j]
			index++
		}
	}

	return result
}

func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
	}
}

func main() {
	// result: [1,2,4,7,5,3,6,8,9]
	// mat := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: [1,2,3,4]
	mat := [][]int{{1,2},{3,4}}

	// result: []
	// mat := [][]int{}

	result := findDiagonalOrder(mat)
	fmt.Printf("result = %v\n", result)
}

