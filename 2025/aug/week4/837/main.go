package main

import (
	"fmt"
)

func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, n + 1)
	dp[0] = 1
	s := float64(1)
	if k <= 0 {
		s = 0
	}

	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}

		if i - maxPts >= 0 && i - maxPts < k {
			s -= dp[i - maxPts]
		}
	}

	result := float64(0)
	for i := k; i <= n; i++ {
		result += dp[i]
	}

	return result
}

func main() {
	// result: 1.00000
	// n := int(10)
	// k := int(1)
	// maxPts := int(10)

	// result: 0.60000
	// n := int(6)
	// k := int(1)
	// maxPts := int(10)

	// result: 0.73278
	n := int(21)
	k := int(17)
	maxPts := int(10)

	// result: 
	// n := int(0)
	// k := int(0)
	// maxPts := int(0)

	result := new21Game(n, k, maxPts)
	fmt.Printf("result = %v\n", result)
}

