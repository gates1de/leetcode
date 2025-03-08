package main
import (
	"fmt"
)

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n % 3 == 2 {
			return false
		}
		n /= 3
	}

	return true
}

func main() {
	// result: true
	// n := int(12)

	// result: true
	// n := int(91)

	// result: false
	n := int(21)

	// result: 
	// n := int(0)

	result := checkPowersOfThree(n)
	fmt.Printf("result = %v\n", result)
}

