package main
import (
	"fmt"
)

func minimumLength(s string) int {
	freq := make([]int, 26)
	result := int(0)
	for _, char := range s {
		freq[char - 'a']++
	}

	for _, count := range freq {
		if count == 0 {
			continue
		}

		if count % 2 == 0 {
			result += 2
		} else {
			result++
		}
	}

	return result
}

func main() {
	// result: 5
	// s := "abaacbcbb"

	// result: 2
	s := "aa"

	// result: 0
	// s := ""

	// result: 0
	// s := ""

	result := minimumLength(s)
	fmt.Printf("result = %v\n", result)
}

