package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2]int, zero + 1)

	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one + 1)
	}
	for i := 0; i <= zero && i <= limit; i++ {
		dp[i][0][0] = 1
	}
	for j := 0; j <= one && j <= limit; j++ {
		dp[0][j][1] = 1
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			if i > limit {
				dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1] - dp[i - limit - 1][j][1]
			} else {
				dp[i][j][0] = dp[i - 1][j][0] + dp[i - 1][j][1]
			}

			dp[i][j][0] = (dp[i][j][0] % modulo + modulo) % modulo

			if j > limit {
				dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0] - dp[i][j - limit - 1][0]
			} else {
				dp[i][j][1] = dp[i][j - 1][1] + dp[i][j - 1][0]
			}

			dp[i][j][1] = (dp[i][j][1] % modulo + modulo) % modulo
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % modulo
}

func main() {
	// result: 2
	// zero := int(1)
	// one := int(1)
	// limit := int(2)

	// result: 1
	// zero := int(1)
	// one := int(2)
	// limit := int(1)

	// result: 14
	zero := int(3)
	one := int(3)
	limit := int(2)

	// result: 
	// zero := int(0)
	// one := int(0)
	// limit := int(0)

	result := numberOfStableArrays(zero, one, limit)
	fmt.Printf("result = %v\n", result)
}

