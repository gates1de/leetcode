package main
import (
	"fmt"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n)
	dp[n - 1] = int64(questions[n - 1][0])

	for i := n - 2; i >= 0; i-- {
		dp[i] = int64(questions[i][0])
		skip := questions[i][1]
		if i + skip + 1 < n {
			dp[i] += dp[i + skip + 1]
		}

		dp[i] = max(dp[i], dp[i + 1])
	}

	return dp[0]
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// questions := [][]int{{3,2},{4,3},{4,4},{2,5}}

	// result: 7
	questions := [][]int{{1,1},{2,2},{3,3},{4,4},{5,5}}

	// result: 
	// questions := [][]int{}

	result := mostPoints(questions)
	fmt.Printf("result = %v\n", result)
}
