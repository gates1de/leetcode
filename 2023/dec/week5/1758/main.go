package main
import (
	"fmt"
)

func minOperations(s string) int {
	n := len(s)
	start := int(0)
	for i, char := range s {
		if i % 2 == 0 {
			if char == '1' {
				start++
			}
		} else {
			if char == '0' {
				start++
			}
		}
	}

	return min(start, n - start)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// s := "0100"

	// result: 0
	// s := "10"

	// result: 2
	s := "1111"

	// result: 
	// s := ""

	result := minOperations(s)
	fmt.Printf("result = %v\n", result)
}

