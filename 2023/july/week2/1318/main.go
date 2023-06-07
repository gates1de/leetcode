package main
import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	result := int(0)
	for a != 0 || b != 0 || c != 0 {
		if (c & 1) == 1 {
			if (a & 1) == 0 && (b & 1) == 0 {
				result++
			}
		} else {
			result += (a & 1) + (b & 1)
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return result
}

func main() {
	// result: 3
	// a := int(2)
	// b := int(6)
	// c := int(5)

	// result: 1
	// a := int(4)
	// b := int(2)
	// c := int(7)

	// result: 0
	a := int(1)
	b := int(2)
	c := int(3)

	// result: 
	// a := int(0)
	// b := int(0)
	// c := int(0)

	result := minFlips(a, b, c)
	fmt.Printf("result = %v\n", result)
}

