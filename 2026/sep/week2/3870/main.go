package main

import (
	"fmt"
)

func countCommas(n int) int {
	result := int(0)
	for place := 1000; place <= n; place *= 1000 {
		result += n - place + 1
	}

	return result
}

func main() {
	// result: 3
	// n := int(1002)

	// result: 0
	n := int(998)

	// result: 0
	// n := int(0)

	result := countCommas(n)
	fmt.Printf("result = %v\n", result)
}
