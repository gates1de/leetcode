package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numOfWays(nums []int) int {
	n := len(nums)
	factorial := make([]int64, n+1)
	inverseFactorial := make([]int64, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * int64(i) % int64(modulo)
	}
	inverseFactorial[n] = modPow(factorial[n], int64(modulo-2))
	for i := n; i > 0; i-- {
		inverseFactorial[i-1] = inverseFactorial[i] * int64(i) % int64(modulo)
	}

	ways := dfs(nums, factorial, inverseFactorial)
	return int((ways + int64(modulo) - 1) % int64(modulo))
}

func dfs(nums []int, factorial []int64, inverseFactorial []int64) int64 {
	n := len(nums)
	if n < 3 {
		return 1
	}

	leftNodes := make([]int, 0, n)
	rightNodes := make([]int, 0, n)
	for i := 1; i < n; i++ {
		num := nums[i]
		if num < nums[0] {
			leftNodes = append(leftNodes, num)
		} else {
			rightNodes = append(rightNodes, num)
		}
	}

	leftWays := dfs(leftNodes, factorial, inverseFactorial)
	rightWays := dfs(rightNodes, factorial, inverseFactorial)
	leftCount := len(leftNodes)
	choose := factorial[n-1] * inverseFactorial[leftCount] % int64(modulo)
	choose = choose * inverseFactorial[n-1-leftCount] % int64(modulo)

	return leftWays * rightWays % int64(modulo) * choose % int64(modulo)
}

func modPow(base int64, exponent int64) int64 {
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
	// result: 1
	// nums := []int{2,1,3}

	// result: 5
	// nums := []int{3,4,5,1,2}

	// result: 0
	nums := []int{1, 2, 3}

	// result:
	// nums := []int{}

	result := numOfWays(nums)
	fmt.Printf("result = %v\n", result)
}
