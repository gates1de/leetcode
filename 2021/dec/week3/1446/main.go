package main
import (
	"fmt"
)

func maxPower(s string) int {
	if len(s) == 1 {
		return 1
	}

	result := int(0)
	count := int(1)
	currentChar := rune(s[0])
	for _, c := range s[1:] {
		if currentChar == c {
			count++
			continue
		}

		if result < count {
			result = count
		}
		currentChar = c
		count = 1
	}

	if result < count {
		result = count
	}

	return result
}

func main() {
	// result: 2
	// s := "leetcode"

	// result: 5
	// s := "abbcccddddeeeeedcba"

	// result: 5
	// s := "triplepillooooow"

	// result: 11
	// s := "hooraaaaaaaaaaay"

	// result: 1
	// s := "tourist"

	// result: 2
	s := "cc"

	// result: 
	// s := ""

	result := maxPower(s)
	fmt.Printf("result = %v\n", result)
}

