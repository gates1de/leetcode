package main
import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func countGoodNumbers(n int64) int {
	return int(quickMultiple(5, (n + 1) / 2) * quickMultiple(4, n / 2) % modulo)
}

func quickMultiple(x, y int64) int64 {
	result := int64(1)
	multiple := x
	for y > 0 {
		if y % 2 == 1 {
			result = result * multiple % modulo
		}

		multiple = multiple * multiple % modulo
		y /= 2
	}

	return result
}

func main() {
	// result: 5
	// n := int64(1)

	// result: 400
	// n := int64(4)

	// result: 564908303
	n := int64(50)

	// result: 
	// n := int64(0)

	result := countGoodNumbers(n)
	fmt.Printf("result = %v\n", result)
}

