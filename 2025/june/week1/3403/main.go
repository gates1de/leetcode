package main
import (
	"fmt"
)

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	n := len(word)
	result := ""
	for i := 0; i < n; i++ {
		result = max(result, word[i:min(i + n - numFriends + 1, n)])
	}

	return result
}

func max(a, b string) string {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: "dbc"
	// word := "dbca"
	// numFriends := int(2)

	// result: "g"
	word := "gggg"
	numFriends := int(4)

	// result: ""
	// word := ""
	// numFriends := int(0)

	result := answerString(word, numFriends)
	fmt.Printf("result = %v\n", result)
}

