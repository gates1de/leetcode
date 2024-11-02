package main
import (
	"fmt"
	"strings"
)

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	if len(words) == 0 {
		return false
	}

	firstWord := words[0]
	lastChar := firstWord[len(firstWord) - 1]

	for _, word := range words[1:] {
		if len(word) == 0{
			continue
		}

		if lastChar != word[0] {
			return false
		}

		lastChar = word[len(word) - 1]
	}

	return lastChar == firstWord[0]
}

func main() {
	// result: true
	// sentence := "leetcode exercises sound delightful"

	// result: true
	// sentence := "eetcode"

	// result: false
	sentence := "Leetcode is cool"

	// result: false
	// sentence := ""

	result := isCircularSentence(sentence)
	fmt.Printf("result = %v\n", result)
}

