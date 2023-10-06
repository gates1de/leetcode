package main
import (
	"fmt"
)

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	memo := make([]int, n + 1)
	return dp(n, &memo)
}

func dp(num int, memo *[]int) int {
	if num <= 3 {
		return num
	}
	if (*memo)[num] != 0 {
		return (*memo)[num]
	}

	result := num
	for i := 2; i < num; i++ {
		result = max(result, i * dp(num - i, memo))
	}

	(*memo)[num] = result
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// n := int(2)

	// result: 36
	n := int(10)

	// result: 
	// n := int(0)

	result := integerBreak(n)
	fmt.Printf("result = %v\n", result)
}

