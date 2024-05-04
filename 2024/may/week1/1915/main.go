package main
import (
	"fmt"
)

func wonderfulSubstrings(word string) int64 {
	freq := make(map[int]int)
	freq[0] = 1

	mask := int(0)
	result := int64(0)

	for _, char := range word {
		bit := int(char - 'a')

		mask ^= (1 << bit)

		result += int64(freq[mask])
		freq[mask]++

		for oddChar := 0; oddChar < 10; oddChar++ {
			result += int64(freq[mask ^ (1 << oddChar)])
		}
	}

	return result
}

func main() {
	// result: 4
	// word := "aba"

	// result: 9
	// word := "aabb"

	// result: 2
	word := "he"

	// result: 
	// word := ""

	result := wonderfulSubstrings(word)
	fmt.Printf("result = %v\n", result)
}

