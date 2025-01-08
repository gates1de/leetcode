package main
import (
	"fmt"
)

func countPrefixSuffixPairs(words []string) int {
	n := len(words)
	result := int(0)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			str1 := words[i]
			str2 := words[j]

			if len(str1) > len(str2) {
				continue
			}

			start := string(str2[:len(str1)])
			end := string(str2[len(str2) - len(str1):])
			if start == str1 && end == str1 {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 4
	// words := []string{"a","aba","ababa","aa"}

	// result: 2
	// words := []string{"pa","papa","ma","mama"}

	// result: 0
	words := []string{"abab","ab"}

	// result: 
	// words := []string{}

	result := countPrefixSuffixPairs(words)
	fmt.Printf("result = %v\n", result)
}

