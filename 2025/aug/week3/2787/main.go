package main

import (
	"fmt"
	"math"
)

const modulo = int(1e9 + 7)

func numberOfWays(n int, x int) int {
	dp := make([]int, n + 1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		val := int(math.Pow(float64(i), float64(x)))
		if val > n {
			break
		}

		for j := n; j >= val; j-- {
			dp[j] = (dp[j] + dp[j - val]) % modulo
		}
	}

	return dp[n]
}

func main() {
	// result: 1
	// n := int(10)
	// x := int(2)

	// result: 2
	n := int(4)
	x := int(1)

	// result: 
	// n := int(0)
	// x := int(0)

	result := numberOfWays(n, x)
	fmt.Printf("result = %v\n", result)
}

