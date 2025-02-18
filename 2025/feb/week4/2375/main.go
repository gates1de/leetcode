package main
import (
	"fmt"
)

func smallestNumber(pattern string) string {
	n := len(pattern)
	result := make([]byte, 0, n)
	stack := make([]int, 0, n)

	for i := 0; i <= n; i++ {
		stack = append(stack, i + 1)

		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				result = append(result, byte(stack[len(stack) - 1]) + 48)
				stack = stack[:len(stack) - 1]
			}
		}
	}

	return string(result)
}

func main() {
	// result: "123549876"
	// pattern := "IIIDIDDD"

	// result: "4321"
	pattern := "DDD"

	// result: ""
	// pattern := ""

	result := smallestNumber(pattern)
	fmt.Printf("result = %v\n", result)
}

