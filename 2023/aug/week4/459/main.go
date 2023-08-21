package main
import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	t := s + s
	fmt.Println(t[1:len(t) - 1])
	if strings.Contains(t[1:len(t) - 1], s) {
		return true
	}
	return false
}

func main() {
	// result: true
	// s := "abab"

	// result: false
	// s := "aba"

	// result: true
	// s := "abcabcabcabc"

	// result: false
	s := "abbbba"

	// result: 
	// s := ""

	result := repeatedSubstringPattern(s)
	fmt.Printf("result = %v\n", result)
}

