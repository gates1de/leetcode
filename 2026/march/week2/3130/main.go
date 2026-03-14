package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][]int, zero + 1)

	for i := range dp {
		dp[i] = make([][]int, one + 1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			for lastBit := 0; lastBit <= 1; lastBit++ {
				if i == 0 {
					if lastBit == 0 || j > limit {
						dp[i][j][lastBit] = 0
					} else {
						dp[i][j][lastBit] = 1
					}
				} else if j == 0 {
					if lastBit == 1 || i > limit {
						dp[i][j][lastBit] = 0
					} else {
						dp[i][j][lastBit] = 1
					}
				} else if lastBit == 0 {
					dp[i][j][lastBit] = dp[i - 1][j][0] + dp[i - 1][j][1]
					if i > limit {
						dp[i][j][lastBit] -= dp[i - limit - 1][j][1]
					}
				} else {
					dp[i][j][lastBit] = dp[i][j - 1][0] + dp[i][j - 1][1]
					if j > limit {
						dp[i][j][lastBit] -= dp[i][j - limit - 1][0]
					}
				}

				dp[i][j][lastBit] %= modulo
				if dp[i][j][lastBit] < 0 {
					dp[i][j][lastBit] += modulo
				}
			}
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

