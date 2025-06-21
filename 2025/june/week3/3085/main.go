package main
import (
	"fmt"
)

func minimumDeletions(word string, k int) int {
	counter := make(map[rune]int)
	for _, char := range word {
		counter[char]++
	}

	result := len(word)
	for _, count1 := range counter {
		deleted := int(0)

		for _, count2 := range counter {
			if count1 > count2 {
				deleted += count2
			} else if count2 > count1 + k {
				deleted += count2 - (count1 + k)
			}
		}

		result = min(result, deleted)
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// word := "aabcaba"
	// k := int(0)

	// result: 2
	// word := "dabdcbdcdcd"
	// k := int(2)

	// result: 1
	word := "aaabaaa"
	k := int(2)

	// result: 
	// word := ""
	// k := int(0)

	result := minimumDeletions(word, k)
	fmt.Printf("result = %v\n", result)
}

