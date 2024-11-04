package main
import (
	"fmt"
)

func compressedString(word string) string {
	if len(word) == 0 {
		return ""
	}

	count := int(1)
	preChar := word[0]
	result := make([]byte, 0, len(word))

	for _, char := range word[1:] {
		if preChar == byte(char) {
			count++
			continue
		}

		for count > 0 {
			if count >= 9 {
				result = append(result, '9')
				result = append(result, preChar)
				count -= 9
			} else {
				result = append(result, '0' + byte(count))
				result = append(result, preChar)
				count = 0
			}
		}

		count = 1
		preChar = byte(char)
	}

	for count > 0 {
		if count >= 9 {
			result = append(result, '9')
			result = append(result, preChar)
			count -= 9
		} else {
			result = append(result, '0' + byte(count))
			result = append(result, preChar)
			count = 0
		}
	}

	return string(result)
}

func main() {
	// result: "1a1b1c1d1e"
	// word := "abcde"

	// result: "9a5a2b"
	word := "aaaaaaaaaaaaaabb"

	// result: ""
	// word := ""

	result := compressedString(word)
	fmt.Printf("result = %v\n", result)
}

