package main
import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	return min(dividend / divisor, math.MaxInt32)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// dividend := int(10)
	// divisor := int(3)

	// dividend := int(7)
	// divisor := int(-3)

	dividend := int(-2147483648)
	divisor := int(-1)

	result := divide(dividend, divisor)
	fmt.Printf("result = %v\n", result)
}

