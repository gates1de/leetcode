package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	result := int(0)
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c := int(math.Sqrt(float64(a * a + b * b + 1)))
			if c <= n && c * c == a * a + b * b {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 2
	// n := int(5)

	// result: 4
	n := int(10)

	// result: 
	// n := int(0)

	result := countTriples(n)
	fmt.Printf("result = %v\n", result)
}

