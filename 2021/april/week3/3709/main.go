package main
import (
	"fmt"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

func main() {
	// result: 1
	// n := int(2)

	// result: 2
	// n := int(3)

	// result: 3
	// n := int(4)

	// result: 832040
	n := int(30)

	result := fib(n)
	fmt.Printf("result = %v\n", result)
}

