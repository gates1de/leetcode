package main
import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	result := dividend / divisor
	if result >= math.MaxInt32 {
		return math.MaxInt32
	} else if result <=  math.MinInt32 {
		return math.MinInt32
	}
	return result
}

func main() {
	// result: 3
	// dividend := int(10)
	// divisor := int(3)

	// result: -2
	// dividend := int(7)
	// divisor := int(-3)

	// result: -2147483647
	dividend := math.MinInt32
	divisor := int(-1)

	// result: 
	// dividend := int(0)
	// divisor := int(0)

	result := divide(dividend, divisor)
	fmt.Printf("result = %v\n", result)
}

