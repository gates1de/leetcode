package main
import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		if len(searchWord) <= len(word) && searchWord == string(word[:len(searchWord)]) {
			return i + 1
		}
	}

	return -1
}

func main() {
	// result: 4
	// sentence := "i love eating burger"
	// searchWord := "burg"

	// result: 2
	// sentence := "this problem is an easy problem"
	// searchWord := "pro"

	// result: -1
	sentence := "i am tired"
	searchWord := "you"

	// result: 
	// sentence := ""
	// searchWord := ""

	result := isPrefixOfWord(sentence, searchWord)
	fmt.Printf("result = %v\n", result)
}

