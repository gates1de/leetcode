package main
import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i, _ := range dp {
		dp[i] = make([]int, len(triangle[i]))
		for j, _ := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	top := triangle[0][0]
	helper(0, 0, top, triangle, dp)
	result := math.MaxInt32
	for _, num := range dp[len(dp) - 1] {
		if result > num {
			result = num
		}
	}
	return result
}

func helper(level int, index int, minSum int, triangle [][]int, dp [][]int) {
	if dp[level][index] <= minSum {
		return
	}
	dp[level][index] = minSum

	level++
	if level == len(triangle) {
		return
	}

	helper(level, index, minSum + triangle[level][index], triangle, dp)
	helper(level, index + 1, minSum + triangle[level][index + 1], triangle, dp)
}

func main() {
	// result: 11
	// triangle := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}

	// result: -10
	// triangle := [][]int{{-10}}

	// result: -1
	triangle := [][]int{{-1},{2,3},{1,-1,-3}}

	// result: 
	// triangle := [][]int{}

	result := minimumTotal(triangle)
	fmt.Printf("result = %v\n", result)
}

