package main
import (
	"fmt"
)

func countVowelStrings(n int) int {
	// vowels := map[string]int{"a": 5, "e": 4, "i": 3, "o": 2, "u": 1}
	// fmt.Printf("vowels = %v\n", vowels)
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
	fmt.Printf("n = %v, charIndex = %v\n", n, charIndex)
	if n == 0 {
		return 1
	}
	result := int(0)
	for j := 0; j < 5 - charIndex; j++ {
		result += count(n - 1, j)
	}
	fmt.Printf("count = %v\n", result)
	return result
}

func main() {
	n := 2
	result := countVowelStrings(n)
	fmt.Printf("result = %v\n", result)
}

