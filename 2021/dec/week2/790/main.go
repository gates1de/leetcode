package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numTilings(n int) int {
	if n <= 2 {
		return n
	}

	full := make([]int, n + 1)
	partial := make([]int, n + 1)
	full[1] = 1
	full[2] = 2
	partial[2] = 1
	for i := 3; i < n + 1; i++ {
		full[i] = (full[i - 1] + full[i - 2] + 2 * partial[i - 1]) % modulo
		partial[i] = (partial[i - 1] + full[i - 2]) % modulo
	}
    return full[n]
}

func main() {
	// result: 5
	// n := int(3)

	// result: 1
	// n := int(1)

	// result: 11
	// n := int(4)

	// result: 979232805
	n := int(1000)

	// result: 
	// n := int(0)

	result := numTilings(n)
	fmt.Printf("result = %v\n", result)
}

