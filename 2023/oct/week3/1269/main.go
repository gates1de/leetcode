package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numWays(steps int, arrLen int) int {
	arrLen = min(arrLen, steps)
	dp := make([]int, arrLen)
	prevDp := make([]int, arrLen)
	prevDp[0] = 1

	for remain := 1; remain <= steps; remain++ {
		dp = make([]int, arrLen)

		for curr := arrLen - 1; curr >= 0; curr-- {
			result := prevDp[curr]

			if curr > 0 {
				result = (result + prevDp[curr - 1]) % modulo
			}

			if curr < arrLen - 1 {
				result = (result + prevDp[curr + 1]) % modulo
			}

			dp[curr] = result
		}

		prevDp = dp
	}

	return dp[0]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// steps := int(3)
	// arrLen := int(2)

	// result: 2
	// steps := int(2)
	// arrLen := int(4)

	// result: 8
	steps := int(4)
	arrLen := int(2)

	// result: 
	// steps := int(0)
	// arrLen := int(0)

	result := numWays(steps, arrLen)
	fmt.Printf("result = %v\n", result)
}

