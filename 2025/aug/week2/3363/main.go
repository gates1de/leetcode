package main

import (
	"fmt"
	"math"
)

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	result := int(0)
	for i := 0; i < n; i++ {
		result += fruits[i][i]
	}

	result += dp(n, fruits)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fruits[i][j], fruits[j][i] = fruits[j][i], fruits[i][j]
		}
	}

	result += dp(n, fruits)
	return result
}

func dp(n int, fruits [][]int) int {
	prev := make([]int, n)
	curr := make([]int, n)
	for i := range prev {
		prev[i] = math.MinInt32
	}

	prev[n - 1] = fruits[0][n - 1]
	for i := 1; i < n - 1; i++ {
		for j := range curr {
			curr[j] = math.MinInt32
		}

		for j := max(n-1-i, i+1); j < n; j++ {
			best := prev[j]
			if j - 1 >= 0 {
				best = max(best, prev[j - 1])
			}

			if j + 1 < n {
				best = max(best, prev[j + 1])
			}

			curr[j] = best + fruits[i][j]
		}

		prev, curr = curr, prev
	}

	return prev[n-1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 100
	// fruits := [][]int{{1,2,3,4},{5,6,8,7},{9,10,11,12},{13,14,15,16}}

	// result: 4
	fruits := [][]int{{1,1},{1,1}}

	// result: 
	// fruits := [][]int{}

	result := maxCollectedFruits(fruits)
	fmt.Printf("result = %v\n", result)
}

