package main
import (
	"fmt"
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		helper(0, i, 0, matrix, dp)
	}

	result := math.MaxInt32
	for _, val := range dp[m - 1] {
		result = min(result, val)
	}

	return result
}

func helper(current int, x int, y int, matrix [][]int, dp [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}

	n := len(matrix[0])
	if n == 0 {
		return
	}

	if x < 0 || y < 0 || x >= n || y >= m {
		return
	}

	current += matrix[y][x]
	memo := dp[y][x]
	if memo != math.MaxInt32 && current >= memo {
		return
	}

	dp[y][x] = current

	if y >= m {
		return
	}

	helper(current, x, y + 1, matrix, dp)
	if x - 1 < n {
		helper(current, x - 1, y + 1, matrix, dp)
	}
	if x + 1 < n {
		helper(current, x + 1, y + 1, matrix, dp)
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 13
	// matrix := [][]int{{2,1,3},{6,5,4},{7,8,9}}

	// result: -59
	matrix := [][]int{{-19,57},{-40,-5}}

	// result: 
	// matrix := [][]int{}

	result := minFallingPathSum(matrix)
	fmt.Printf("result = %v\n", result)
}

