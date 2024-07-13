package main
import (
	"fmt"
)

func findTheWinner(n int, k int) int {
	return helper(n, k) + 1
}

func helper(n int, k int) int {
	if n == 1 {
		return 0
	}

	return (helper(n - 1, k) + k) % n
}

func main() {
	// result: 3
	// n := int(5)
	// k := int(2)

	// result: 1
	n := int(6)
	k := int(5)

	// result: 
	// n := int(0)
	// k := int(0)

	result := findTheWinner(n, k)
	fmt.Printf("result = %v\n", result)
}

