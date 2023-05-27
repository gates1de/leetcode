package main
import (
	"fmt"
)

func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][][]int, 2)
	for p := 0; p < 2; p++ {
		dp[p] = make([][]int, n + 1)
		for i := 0; i <= n; i++ {
			dp[p][i] = make([]int, n + 1)
			for j := 0; j <= n; j++ {
				dp[p][i][j] = -1
			}
		}
	}
	return helper(0, 0, 1, piles, dp)
}

func helper(p int, i int, m int, piles []int, dp [][][]int) int {
	n := len(piles)
	if i == n {
		return 0
	}

	if dp[p][i][m] != -1 {
		return dp[p][i][m]
	}

	result := int(-1)
	s := int(0)
	if p == 1 {
		result = 1000000
	}

	for x := 1; x <= min(2 * m, n - i); x++ {
		s += piles[i + x - 1]

		if p == 0 {
			result = max(result, s + helper(1, i + x, max(m, x), piles, dp))
		} else {
			result = min(result, helper(0, i + x, max(m, x), piles, dp))
		}
	}

    dp[p][i][m] = result
	return result
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
	// result: 10
	// piles := []int{2,7,9,4,4}

	// result: 104
	piles := []int{1,2,3,4,5,100}

	// result: 
	// piles := []int{}

	result := stoneGameII(piles)
	fmt.Printf("result = %v\n", result)
}

