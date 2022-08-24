package main
import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	m := map[int]bool{}
	for i := 1; i < math.MaxInt32; i *= 3 {
		m[i] = true
	}
	return m[n]
}

func main() {
	// result: true
	// n := int(27)

	// result: false
	// n := int(0)

	// result: true
	// n := int(9)

	// result: false
	// n := math.MaxInt32 - 1

	// result: true
	n := 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3

	// result: 
	// n := int(0)

	result := isPowerOfThree(n)
	fmt.Printf("result = %v\n", result)
}

