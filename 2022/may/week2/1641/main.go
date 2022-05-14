package main
import (
	"fmt"
)

func countVowelStrings(n int) int {
	if n == 0 {
		return 0
	}

	return count(n, 0)
}

// charIndex is
// a = 0, e = 1, i = 2, o = 3, u = 4
func count(n int, charIndex int) int {
	if n == 0 {
		return 1
	}

	result := int(0)
	for j := 0; j < 5 - charIndex; j++ {
		result += count(n - 1, 4 - j)
	}

	return result
}

func main() {
	// result: 5
	// n := int(1)

	// result: 15
	// n := int(2)

	// result: 35
	// n := int(3)

	// result: 70
	// n := int(4)

	// result: 126
	// n := int(5)

	// result: 210
	// n := int(6)

	// result: 66045
	n := int(33)

	// result: 
	// n := int(0)

	result := countVowelStrings(n)
	fmt.Printf("result = %v\n", result)
}

