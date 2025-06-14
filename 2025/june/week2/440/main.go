package main
import (
	"fmt"
)

func findKthNumber(n int, k int) int {
	num := int(1)
	k--

	for k > 0 {
		step := countSteps(n, num, num + 1)
		if step <= k {
			num++
			k -= step
		} else {
			num *= 10
			k--
		}
	}

	return num
}

func countSteps(n int, prefix1 int, prefix2 int) int {
	result := int(0)

	for prefix1 <= n {
		result += min(n + 1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 10
	// n := int(13)
	// k := int(2)

	// result: 1
	n := int(1)
	k := int(1)

	// result: 
	// n := int(0)
	// k := int(0)

	result := findKthNumber(n, k)
	fmt.Printf("result = %v\n", result)
}

