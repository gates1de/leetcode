package main
import (
	"fmt"
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	result := math.MaxInt32
	memo := make([][]int, n)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo {
			memo[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		helper(i, 0, 0, matrix, memo, &result)
	}
	return result
}

func helper(x int, y int, sum int, matrix [][]int, memo [][]int, result *int) {
	n := len(matrix)
	if n == 0 || x < 0 || n <= x || y < 0 || n <= y {
		return
	}

	sum += matrix[y][x]
	if y == n - 1 {
		if sum < *result {
			*result = sum
		}
		return
	}

	if sum >= memo[y][x] {
		return
	}
	memo[y][x] = sum

	helper(x - 1, y + 1, sum, matrix, memo, result)
	helper(x, y + 1, sum, matrix, memo, result)
	helper(x + 1, y + 1, sum, matrix, memo, result)
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

