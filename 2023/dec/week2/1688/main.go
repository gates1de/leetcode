package main
import (
	"fmt"
)

func numberOfMatches(n int) int {
	result := int(0)

	for num := n; num >= 2; num /= 2 {
		if num % 2 == 0 {
			result += num / 2
		} else {
			result += (num - 1) / 2 + 1
		}
	}

	return result
}

func main() {
	// result: 6
	// n := int(7)

	// result: 13
	n := int(14)

	// result: 
	// n := int(0)

	result := numberOfMatches(n)
	fmt.Printf("result = %v\n", result)
}

