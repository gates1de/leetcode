package main
import (
	"fmt"
)

func maximumLength(s string) int {
	counter := make(map[string]int)

	for start, _ := range s {
		current := make([]byte, 0, len(s))

		for end := start; end < len(s); end++ {
			if len(current) == 0 || current[len(current) - 1] == s[end] {
				current = append(current, s[end])
				counter[string(current)]++
			} else {
				break
			}
		}
	}

	result := int(0)
	for key, count := range counter {
		if count >= 3 && len(key) > result {
			result = len(key)
		}
	}

	if result == 0 {
		return -1
	}

	return result
}

func main() {
	// result: 2
	// s := "aaaa"

	// result: -1
	// s := "abcdef"

	// result: 1
	s := "abcaba"

	// result: 
	// s := ""

	result := maximumLength(s)
	fmt.Printf("result = %v\n", result)
}

