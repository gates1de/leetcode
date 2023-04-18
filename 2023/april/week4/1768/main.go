package main
import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	bytes := make([]byte, len(word1) + len(word2))
	i := int(0)
	j := int(0)
	for k, _ := range bytes {
		if (k % 2 == 0 && i < len(word1)) || j >= len(word2) {
			bytes[k] = word1[i]
			i++
		} else {
			bytes[k] = word2[j]
			j++
		}
	}
	return string(bytes)
}

func main() {
	// result: "apbqcr"
	// word1 := "abc"
	// word2 := "pqr"

	// result: "apbqrs"
	// word1 := "ab"
	// word2 := "pqrs"

	// result: "apbqcd"
	word1 := "abcd"
	word2 := "pq"

	// result: 
	// word1 := ""
	// word2 := ""

	result := mergeAlternately(word1, word2)
	fmt.Printf("result = %v\n", result)
}

