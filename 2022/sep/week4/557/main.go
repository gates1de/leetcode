package main
import (
	"fmt"
)

func reverseWords(s string) string {
	startIndex := int(0)
	result := ""
	s += " "
	for i, char := range s {
		if char != rune(' ') {
			continue
		}

		reversedChars := make([]byte, i - startIndex)
		index := int(0)
		for j := i - 1; j >= startIndex; j-- {
			reversedChars[index] = s[j]
			index++
		}

		result += string(reversedChars) + " "
		startIndex = i + 1
	}

	return result[:len(result) - 1]
}

func main() {
	// result: "s'teL ekat edoCteeL tsetnoc"
	// s := "Let's take LeetCode contest"

	// result: "doG gniD"
	// s := "God Ding"

	// result: "d"
	s := "d"

	// result: 
	// s := ""

	result := reverseWords(s)
	fmt.Printf("result = %v\n", result)
}

