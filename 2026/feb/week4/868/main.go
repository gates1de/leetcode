package main

import (
	"fmt"
)

func binaryGap(n int) int {
	a := make([]int, 32)
	t := int(0)
	for i := range a {
		if ((n >> i) & 1) != 0 {
			a[t] = i
			t++
		}
	}

	result := int(0)
	for i := range t - 1 {
		result = max(result, a[i + 1] - a[i])
	}

	return result
}

func main() {
	// result: 2
	// n := int(22)

	// result: 0
	// n := int(8)

	// result: 2
	n := int(5)

	// result: 
	// n := int(0)

	result := binaryGap(n)
	fmt.Printf("result = %v\n", result)
}
