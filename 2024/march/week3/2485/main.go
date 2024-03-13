package main
import (
	"fmt"
)

func pivotInteger(n int) int {
	left := int(1)
	right := n

	totalSum := n * (n + 1) / 2

	for left < right {
		mid := (left + right) / 2

		if mid * mid - totalSum < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left * left - totalSum == 0 {
		return left
	}

	return -1
}

func main() {
	// result: 6
	// n := int(8)

	// result: 1
	// n := int(1)

	// result: -1
	n := int(4)

	// result: 
	// n := int(0)

	result := pivotInteger(n)
	fmt.Printf("result = %v\n", result)
}

