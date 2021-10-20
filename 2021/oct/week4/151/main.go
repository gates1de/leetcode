package main
import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	reg := regexp.MustCompile(`\ +`)
	words := reg.Split(s, -1)
	reverse(words)
	return strings.Join(words, " ")
}

func reverse(arr []string) {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - i - 1] = arr[len(arr) - i - 1], arr[i]
	}
}

func main() {
	// result: "blue is sky the"
	// s := "the sky is blue"

	// result: "world hello"
	// s := "  hello world  "

	// result: "example good a"
	// s := "a good   example"

	// result: "Alice Loves Bob"
	s := "  Bob    Loves  Alice   "

	// result: 
	// s := ""

	result := reverseWords(s)
	fmt.Printf("result = %v\n", result)
}

