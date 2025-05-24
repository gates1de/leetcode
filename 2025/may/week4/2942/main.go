package main
import (
	"fmt"
)

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0, len(words))

	for i, word := range words {
		for _, char := range word {
			if byte(char) == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func main() {
	// result: [0,1]
	// words := []string{"leet","code"}
	// x := byte('e')

	// result: [0,2]
	// words := []string{"abc","bcd","aaaa","cbc"}
	// x := byte('a')

	// result: []
	words := []string{}
	x := byte('z')

	// result: []
	// words := []string{}
	// x := byte('')

	result := findWordsContaining(words, x)
	fmt.Printf("result = %v\n", result)
}

