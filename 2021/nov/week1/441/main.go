package main
import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	result := int(1)
	for n > 1 {
		n -= result + 1
		// fmt.Printf("n = %v\n", n)
		if n <= 0 {
			break
		}
		result++
	}

	return result
}

func main() {
	// result: 2
	// n := int(5)

	// result: 3
	// n := int(8)

	// result: 1
	// n := int(1)

	// result: 1
	// n := int(2)

	// result: 65535
	n := int(math.MaxInt32 - 1)

	// result: 
	// n := int(0)

	// result: 
	// n := int(0)

	result := arrangeCoins(n)
	fmt.Printf("result = %v\n", result)
}

