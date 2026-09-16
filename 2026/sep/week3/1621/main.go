package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numberOfSets(n int, k int) int {
	limit := n + k - 1
	factorial := make([]int64, limit+1)
	inverseFactorial := make([]int64, limit+1)
	factorial[0] = 1
	for i := 1; i <= limit; i++ {
		factorial[i] = factorial[i-1] * int64(i) % int64(modulo)
	}

	inverseFactorial[limit] = modPow(factorial[limit], int64(modulo-2))
	for i := limit; i >= 1; i-- {
		inverseFactorial[i-1] = inverseFactorial[i] * int64(i) % int64(modulo)
	}

	result := factorial[limit]
	result = result * inverseFactorial[2*k] % int64(modulo)
	result = result * inverseFactorial[limit-2*k] % int64(modulo)
	return int(result)
}

func modPow(base, exponent int64) int64 {
	result := int64(1)
	mod := int64(modulo)
	for exponent > 0 {
		if exponent&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exponent >>= 1
	}
	return result
}

func main() {
	// result: 5
	// n := int(4)
	// k := int(2)

	// result: 3
	// n := int(3)
	// k := int(1)

	// result: 796297179
	n := int(30)
	k := int(7)

	// result: 0
	// n := int(0)
	// k := int(0)

	result := numberOfSets(n, k)
	fmt.Printf("result = %v\n", result)
}
