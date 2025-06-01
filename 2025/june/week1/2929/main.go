package main
import (
	"fmt"
)

func distributeCandies(n int, limit int) int64 {
	result := int64(0)
	for i := 0; i <= min(limit, n); i++ {
		if n - i > 2 * limit {
			continue
		}

		result += int64(min(n - i, limit) - max(0, n - i - limit) + 1)
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// n := int(5)
	// limit := int(2)

	// result: 10
	n := int(3)
	limit := int(3)

	// result: 
	// n := int(0)
	// limit := int(0)

	result := distributeCandies(n, limit)
	fmt.Printf("result = %v\n", result)
}

