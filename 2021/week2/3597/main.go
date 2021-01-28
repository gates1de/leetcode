package main
import (
	"fmt"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := concat(word1)
	s2 := concat(word2)
	return s1 == s2
}

func concat(list []string) string {
	result := ""
	for _, s := range list {
		result += s
	}
	return result
}

func main() {
	word1 := []string{"ac", "b"}
	word2 := []string{"ab", "c"}
	result := arrayStringsAreEqual(word1, word2)
	fmt.Printf("result = %v\n", result)
}

