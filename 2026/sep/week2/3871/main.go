package main

import (
	"fmt"
)

func countCommas(n int64) int64 {
	result := int64(0)
	for place := int64(1000); place <= n; place *= 1000 {
		result += n - place + 1
	}

	return result
}

func main() {
	// result: 3
	// n := int64(1002)

	// result: 0
	n := int64(998)

	// result: 0
	// n := int64(0)

	result := countCommas(n)
	fmt.Printf("result = %v\n", result)
}
