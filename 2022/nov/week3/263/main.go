package main
import (
	"fmt"
)

func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	for true {
		if n % 2 == 0 {
			n /= 2
		}
		if n % 3 == 0 {
			n /= 3
		}
		if n % 5 == 0 {
			n /= 5
		}

		if n == 1 {
			return true
		}

		if n % 2 != 0 && n % 3 != 0 && n % 5 != 0 {
			return false
		}
	}

	return false
}

func main() {
	// result: true
	// n := int(6)

	// result: true
	// n := int(1)

	// result: false
	n := int(14)

	// result: 
	// n := int(0)

	result := isUgly(n)
	fmt.Printf("result = %v\n", result)
}

