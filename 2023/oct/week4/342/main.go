package main
import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {
    m := make(map[int]bool)
    for i := 1; i < math.MaxInt32; i *= 4 {
        m[i] = true
    }
    return m[n]
}

func main() {
	// result: true
	// n := int(16)

	// result: false
	// n := int(5)

	// result: true
	n := int(1)

	// result: 
	// n := int(0)

	result := isPowerOfFour(n)
	fmt.Printf("result = %v\n", result)
}

