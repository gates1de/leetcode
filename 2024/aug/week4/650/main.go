package main
import (
	"fmt"
)

func minSteps(n int) int {
	result := int(0)
	d := int(2)

	for n > 1 {
		for n % d == 0 {
			result += d
			n /= d
		}

		d++
	}

	return result
}

func main() {
	// result: 3
	// n := int(3)

	// result: 0
	n := int(1)

	// result: 
	// n := int(0)

	result := minSteps(n)
	fmt.Printf("result = %v\n", result)
}

