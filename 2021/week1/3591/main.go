package main
import (
	"fmt"
)

func countArrangement(n int) int {
	result := 0
	if n <= 3 {
		return n
	}

	// perm := make([]int, n)
	// for i := 1; i < n; i++ {
	// 	for j := 1; j < n; j++ {
	// 	}
	// }
	// fmt.Printf("perm = %v\n", perm)
	return result
}

func main() {
	result := countArrangement(1)
	fmt.Printf("result = %v\n", result)
}

