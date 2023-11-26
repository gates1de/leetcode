package main
import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	heights := make([]int, n)
	result := int(0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}

		heightsCopy := make([]int, n)
		copy(heightsCopy, heights)
		sort.Ints(heightsCopy)

		for i := 0; i < n; i++ {
			result = max(result, heightsCopy[i] * (n - i))
		}
	}

	return result
}

// Time Limit Exceeded
func ngSolution(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	prevHeights := make([][2]int, 0)
	result := int(0)

	for row := 0; row < m; row++ {
		heights := make([][2]int, 0, m * n)
		seen := make([]bool, n)

		for _, pair := range prevHeights {
			height := pair[0]
			col := pair[1]

			if matrix[row][col] == 1 {
				heights = append(heights, [2]int{height + 1, col})
				seen[col] = true
			}
		}

		for col := 0; col < n; col++ {
			if !seen[col] && matrix[row][col] == 1 {
				heights = append(heights, [2]int{1, col})
			}
		}

		for i, height := range heights {
			result = max(result, height[0] * (i + 1))
		}

		prevHeights = heights
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// matrix := [][]int{{0,0,1},{1,1,1},{1,0,1}}

	// result: 3
	// matrix := [][]int{{1,0,1,0,1}}

	// result: 2
	matrix := [][]int{{1,1,0},{1,0,1}}

	// result: 
	// matrix := [][]int{}

	result := largestSubmatrix(matrix)
	fmt.Printf("result = %v\n", result)
}

