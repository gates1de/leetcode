package main
import (
	"fmt"
)

func longestIdealString(s string, k int) int {
	n := len(s)
	dp := make([]int, 26)
	result := int(0)

	for i := 0; i < n; i++ {
		current := int(s[i] - 'a')
		best := int(0)

		for prev := 0; prev < 26; prev++ {
			if abs(prev - current) <= k {
				best = max(best, dp[prev])
			}
		}

		dp[current] = max(dp[current], best + 1)
		result = max(result, dp[current])
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}


func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// s := "acfgbd"
	// k := int(2)

	// result: 4
	s := "abcd"
	k := int(3)

	// result: 
	// s := ""
	// k := int(0)

	result := longestIdealString(s, k)
	fmt.Printf("result = %v\n", result)
}

