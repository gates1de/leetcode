package main
import (
	"fmt"
)

func minEnd(n int, x int) int64 {
	result := int64(x)
	n--

	for mask := int64(1); n > 0; mask <<= 1 {
		if (mask & int64(x)) == 0 {
			result |= (int64(n) & 1) * mask
			n >>= 1
		}
	}

	return result
}

func main() {
	// result: 6
	// n := int(3)
	// x := int(4)

	// result: 15
	n := int(2)
	x := int(7)

	// result: 
	// n := int(0)
	// x := int(0)

	result := minEnd(n, x)
	fmt.Printf("result = %v\n", result)
}

