package main
import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	round := int(0)
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, m * n)
	index := int(0)
	maxRound := min(m, n) / 2
	if min(m, n) % 2 == 0 {
		maxRound--
	}
	// fmt.Printf("maxRound = %v\n", maxRound)
	for round <= maxRound {
		for i := round; i < n - round; i++ {
			result[index] = matrix[round][i]
			index++
		}
		for i := round + 1; i < m - round; i++ {
			result[index] = matrix[i][n - 1 - round]
			index++
		}

		if index >= len(result) {
			break
		}

		for i := n - 1 - round - 1; i > round; i-- {
			result[index] = matrix[m - 1 - round][i]
			index++
		}
		for i := m - 1 - round; i > round; i-- {
			result[index] = matrix[i][round]
			index++
		}
		round++
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [1,2,3,6,9,8,7,4,5]
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// result: [1,2,3,4,8,12,11,10,9,5,6,7]
	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// }

	// result: [1,2,3,4,5,6,-12,12,15,-15,14,-14,13,-13,7,-7,-8,-9,-10,-11,11,10,9,8]
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6},
		// {-7, -8, -9, -10, -11, -12},
		{7, 8, 9, 10, 11, 12},
		{-13, 13, -14, 14, -15, 15},
	}

	// result: [1]
	// matrix := [][]int{{1}}

	// result: 
	// matrix := [][]int{}

	result := spiralOrder(matrix)
	fmt.Printf("result = %v\n", result)
}

