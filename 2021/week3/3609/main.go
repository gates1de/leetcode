package main
import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	chars := strings.Split(s, "")
	charsLen := len(chars)
	if charsLen == 0 || charsLen == 1 {
		return s
	}
	palindrome := chars[:1]
	for i, _ := range chars {
		substring := substring(chars[i:], len(palindrome))
		// fmt.Printf("substring = %v\n", substring)
		if len(palindrome) < len(substring) && isPalindrome(substring) {
			palindrome = substring
		}
	}

	return strings.Join(palindrome, "")
}

func substring(chars []string, currentPalLen int) []string {
	charsLen := len(chars)
	if charsLen == 0 || charsLen == 1 {
		return chars
	}
	target := chars[0]
	// fmt.Printf("target = %v, chars = %v\n", target, chars)
	for i := len(chars) - 1; i > 0; i-- {
		c := chars[i]
		if target != c {
			continue
		}

		substring := chars[:i + 1]
		if currentPalLen > len(substring) {
			continue
		}
		if isPalindrome(substring) {
			return substring
		}
	}
	return chars[:1]
}

func isPalindrome(chars []string) bool {
	if len(chars) == 1 {
		return true
	}

	i := int(0)
	j := int(len(chars) - 1)
	for i <= j {
		left := chars[i]
		right := chars[j]
		if left != right {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	// s := "babad"
	// s := "cbbd"
	// s := "a"
	// s := "ac"
	// s := "ccc"
	// s := "chictac"
	// s := "aaaa"
    s := "aacabdkacaa"
	result := longestPalindrome(s)
	fmt.Printf("result = %v\n", result)
}

