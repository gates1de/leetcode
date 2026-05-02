package main

import (
	"fmt"
)

func rotatedDigits(n int) int {
	dp := make([]int, n + 1)
	result := int(0)

	for i := range n + 1 {
		if i < 10 {
			if i == 0 || i == 1 || i == 8 {
				dp[i] = 1
			} else if i == 2 || i == 5 || i == 6 || i == 9 {
				dp[i] = 2
				result++;
			} else {
				dp[i] = 0
			}
		} else {
			a := dp[i / 10]
			b := dp[i % 10]

			if a == 1 && b == 1 {
				dp[i] = 1
			} else if a >= 1 && b >= 1 {
				dp[i] = 2
				result++
			} else {
				dp[i] = 0
			}
		}
	}

	return result
}

func main() {
	// result: 0
	// n := int(1)

	// result: 1
	// n := int(2)

	// result: 4
	n := int(10)

	// result: 
	// n := int(0)

	result := rotatedDigits(n)
	fmt.Printf("result = %v\n", result)
}

