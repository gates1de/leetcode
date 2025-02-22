package main
import (
	"fmt"
)

func numTilePossibilities(tiles string) int {
	charCount := make([]int, 26)
	for _, char := range tiles {
		charCount[char - 'A']++
	}
	return findSequences(charCount)
}

func findSequences(charCount []int) int {
	result := int(0)

	for char := 0; char < 26; char++ {
		if charCount[char] == 0 {
			continue
		}

		result++
		charCount[char]--
		result += findSequences(charCount)
		charCount[char]++
	}

	return result
}

func main() {
	// result: 8
	// tiles := "AAB"

	// result: 188
	// tiles := "AAABBC"

	// result: 1
	tiles := "V"

	// result: 
	// tiles := ""

	result := numTilePossibilities(tiles)
	fmt.Printf("result = %v\n", result)
}

