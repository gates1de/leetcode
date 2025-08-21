package main

import (
	"fmt"
)

func numSubmat(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return -1
	}

	n := len(mat[0])
	heights := make([]int, n)
	result := int(0)

	for _, row := range mat {
		for i := 0; i < n; i++ {
			if row[i] == 0 {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}

		stack := [][3]int{{-1, 0, -1}}
		for i, h := range heights {
			for len(stack) > 1 && stack[len(stack) - 1][2] >= h {
				stack = stack[:len(stack) - 1]
			}

			top := stack[len(stack) - 1]
			j := top[0]
			prev := top[1]
			current := prev + (i - j) * h
			stack = append(stack, [3]int{i, current, h})
			result += current
		}
	}

	return result
}

func main() {
	// result: 13
	// mat := [][]int{{1,0,1},{1,1,0},{1,1,0}}

	// result: 24
	mat := [][]int{{0,1,1,0},{0,1,1,1},{1,1,1,0}}

	// result: 
	// mat := [][]int{}

	result := numSubmat(mat)
	fmt.Printf("result = %v\n", result)
}

