package main
import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	if m < n {
		return combination(m + n - 2, m - 1)
	}
	return combination(m + n - 2, n - 1)
}

func combination(n int, r int) int {
	result := int(1)
	for i := n; i > n - r; i-- {
		result *= i
	}
	for i := 1; i <= r; i++ {
		result /= i
	}
	return result
}

func main() {
	// result: 28
	// m := int(3)
	// n := int(7)

	// result: 3
	// m := int(3)
	// n := int(2)

	// result: 6
	// m := int(3)
	// n := int(3)

	// result: 1
	// m := int(1)
	// n := int(1)

	// result: 193536720
	m := int(23)
	n := int(12)

	// result: 
	// m := int(0)
	// n := int(0)

	result := uniquePaths(m, n)
	fmt.Printf("result = %v\n", result)
}

