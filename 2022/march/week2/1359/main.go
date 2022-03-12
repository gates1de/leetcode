package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countOrders(n int) int {
	if n == 1 {
		return 1
	}

	result := int(1)
	for i := 2; i <= n; i++ {
		result *= combination(i * 2, 2)
		result %= modulo
	}
	return result
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
	// result: 1
	// n := int(1)

	// result: 6
	// n := int(2)

	// result: 90
	// n := int(3)

	// result: 2520
	// n := int(4)

	// result: 764678010
	n := int(500)

	// result: 
	// n := int(0)

	result := countOrders(n)
	fmt.Printf("result = %v\n", result)
}

