package main
import (
	"fmt"
)

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	dp := make([][]int, n + 1)
	for i, _ := range dp {
		dp[i] = make([]int, k + 1)
	}

	for i := 1; i <= n; i++ {
		for coins := 0; coins <= k; coins++ {
			currentSum := int(0)
			for currentCoins := 0; currentCoins <= min(len(piles[i - 1]), coins); currentCoins++ {
				if currentCoins > 0 {
					currentSum += piles[i - 1][currentCoins - 1]
				}
				dp[i][coins] = max(dp[i][coins], dp[i - 1][coins - currentCoins] + currentSum)
			}
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 101
	// piles := [][]int{{1,100,3},{7,8,9}}
	// k := int(2)

	// result: 706
	piles := [][]int{{100},{100},{100},{100},{100},{100},{1,1,1,1,1,1,700}}
	k := int(7)

	// result: 
	// piles := [][]int{}
	// k := int(0)

	result := maxValueOfCoins(piles, k)
	fmt.Printf("result = %v\n", result)
}

