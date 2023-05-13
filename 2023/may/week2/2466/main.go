package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high + 1)
	dp[0] = 1
	for end := 1; end <= high; end++ {
		if end >= zero {
			dp[end] += dp[end - zero]
		}
		if end >= one {
			dp[end] += dp[end - one]
		}
		dp[end] %= modulo
	}

	result := int(0)
	for i := low; i <= high; i++ {
		result += dp[i]
		result %= modulo
	}

	return result
}

func main() {
	// result: 8
	// low := int(3)
	// high := int(3)
	// zero := int(1)
	// one := int(1)

	// result: 5
	low := int(2)
	high := int(3)
	zero := int(1)
	one := int(2)

	// result: 
	// low := int(0)
	// high := int(0)
	// zero := int(0)
	// one := int(0)

	result := countGoodStrings(low, high, zero, one)
	fmt.Printf("result = %v\n", result)
}

