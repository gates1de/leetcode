package main
import (
	"fmt"
)

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	symbol := int(1)

	for row := n; row > 1; row-- {
		totalElements := pow(2, row - 1)
		halfElements := totalElements / 2

		if k > halfElements {
			symbol = 1 - symbol
			k -= halfElements
		}
	}

	if symbol != 0 {
		return 0
	}

	return 1
}

func pow(a, b int) int {
	result := int(1)
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	// result: 0
	// n := int(1)
	// k := int(1)

	// result: 0
	// n := int(2)
	// k := int(1)

	// result: 1
	n := int(2)
	k := int(2)

	// result: 
	// n := int(0)
	// k := int(0)

	result := kthGrammar(n, k)
	fmt.Printf("result = %v\n", result)
}

