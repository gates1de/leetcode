package main
import (
	"fmt"
)

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	chars := []byte(palindrome)

	if n == 1 {
		return ""
	} else if n == 2 {
		if chars[0] == byte('a') {
			chars[1] = byte('b')
		} else {
			chars[0] = byte('a')
		}
		return string(chars)
	} else if n == 3 {
		if chars[0] == byte('a') {
			chars[2] = byte('b')
		} else {
			chars[0] = byte('a')
		}
		return string(chars)
	}

	for i := 0; i < n / 2; i++ {
		left := chars[i]
		rightIndex := n - 1 - i

		if i == rightIndex {
			continue
		}

		if left == byte('a') {
			if chars[i + 1] != byte('a') {
				continue
			}
			chars[rightIndex] = byte('b')
		} else {
			chars[i] = byte('a')
		}
		return string(chars)
	}

	return ""
}

func main() {
	// result: "aaccba"
	// palindrome := "abccba"

	// result: ""
	// palindrome := "a"

	// result: "aacca"
	// palindrome := "accca"

	// result: "acb"
	// palindrome := "aca"

	// result: "ab"
	// palindrome := "aa"

	// result: "aab"
	// palindrome := "bab"

	// result: "ab"
	// palindrome := "bb"

	// result: "aaab"
	palindrome := "aaaa"

	// result: 
	// palindrome := ""

	result := breakPalindrome(palindrome)
	fmt.Printf("result = %v\n", result)
}

