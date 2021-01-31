package main
import (
	"fmt"
)

func countVowelStrings(n int) int {
	result := int(0)

	if n == 0 {
		return 0
	}
	result = count(n, 0)
	return result
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
	// n := 1
	// n := 2
	n := 33
	result := countVowelStrings(n)
	fmt.Printf("result = %v\n", result)
}

