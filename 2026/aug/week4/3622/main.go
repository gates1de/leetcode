package main

import (
	"fmt"
)

func checkDivisibility(n int) bool {
	sum, product := 0, 1
	for value := n; value > 0; value /= 10 {
		digit := value % 10
		sum += digit
		product *= digit
	}

	return n % (sum + product) == 0
}

func main() {
	// result: true
	// n := int(99)

	// result: false
	n := int(23)

	// result:
	// n := int(0)

	result := checkDivisibility(n)
	fmt.Printf("result = %v\n", result)
}
