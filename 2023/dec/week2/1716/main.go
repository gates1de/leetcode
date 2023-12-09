package main
import (
	"fmt"
)

func totalMoney(n int) int {
	result := int(0)

	week := int(1)
	for n > 7 {
		result += 28 + (week - 1) * 7

		if n - 7 <= 0 {
			break
		}
		n -= 7
		week++
	}

	for i := 1; i <= n; i++ {
		result += (week - 1) + i
	}

	return result
}

func main() {
	// result: 10
	// n := int(4)

	// result: 37
	// n := int(10)

	// result: 96
	n := int(20)

	// result: 
	// n := int(0)

	result := totalMoney(n)
	fmt.Printf("result = %v\n", result)
}

