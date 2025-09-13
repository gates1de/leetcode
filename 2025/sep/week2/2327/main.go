package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func peopleAwareOfSecret(n int, delay int, forget int) int {
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i % forget] = (modulo + dp[(i + forget - delay) % forget] - dp[i % forget]) % modulo
			continue
		}

		dp[i % forget] = (modulo + dp[(i + forget - delay) % forget] - dp[i % forget] + dp[(i - 1) % forget]) % modulo
	}

	return sum(dp) % modulo
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 5
	// n := int(6)
	// delay := int(2)
	// forget := int(4)

	// result: 6
	n := int(4)
	delay := int(1)
	forget := int(3)

	// result: 
	// n := int(0)
	// delay := int(0)
	// forget := int(0)

	result := peopleAwareOfSecret(n, delay, forget)
	fmt.Printf("result = %v\n", result)
}

