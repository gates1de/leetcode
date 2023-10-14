package main
import (
	"fmt"
	"math"
)

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n + 1)
	prevDp := make([]int, n + 1)
	for i, _ := range prevDp {
		prevDp[i] = math.MaxInt32
	}
	prevDp[0] = 0

	for i := n - 1; i >= 0; i-- {
		dp = make([]int, n + 1)
		for remain := 1; remain <= n; remain++ {
			paint := cost[i] + prevDp[max(0, remain - 1 - time[i])]
			dontPaint := prevDp[remain]
			dp[remain] = min(paint, dontPaint)
		}

		prevDp = dp
	}

	return dp[n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// cost := []int{1,2,3,2}
	// time := []int{1,2,3,2}

	// result: 4
	cost := []int{2,3,4,2}
	time := []int{1,1,1,1}

	// result: 
	// cost := []int{}
	// time := []int{}

	result := paintWalls(cost, time)
	fmt.Printf("result = %v\n", result)
}

