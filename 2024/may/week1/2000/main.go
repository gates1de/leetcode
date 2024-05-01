package main
import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {
	for i, char := range word {
		if byte(char) == ch {
			return reverse(string(word[:i + 1])) + string(word[i + 1:])
		}
	}

	return word
}

func reverse(word string) string {
	i := int(0)
	j := len(word) - 1
	chars := []byte(word)

	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}

	return string(chars)
}

func main() {
	// result: "dcbaefd"
	// word := "abcdefd"
	// ch := byte('d')

	// result: "zxyxxe"
	// word := "xyxzxe"
	// ch := byte('z')

	// result: "abcd"
	word := "abcd"
	ch := byte('z')

	// result: ""
	// word := ""
	// ch := ''

	result := reversePrefix(word, ch)
	fmt.Printf("result = %v\n", result)
}

