package main
import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int)
	result := make([]string, 0, len(m))

	for _, word := range strings.Split(s1, " ") {
		m[word]++
	}
	for _, word := range strings.Split(s2, " ") {
		m[word]++
	}

	for word, count := range m {
		if count == 1 {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	// result: ["sweet","sour"]
	// s1 := "this apple is sweet"
	// s2 := "this apple is sour"

	// result: ["banana"]
	s1 := "apple apple"
	s2 := "banana"

	// result: []
	// s1 := ""
	// s2 := ""

	result := uncommonFromSentences(s1, s2)
	fmt.Printf("result = %v\n", result)
}

