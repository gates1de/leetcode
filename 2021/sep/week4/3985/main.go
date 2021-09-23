package main
import (
	"fmt"
)

func breakPalindrome(palindrome string) string {
	for i := 0; i < len(palindrome); i++ {
		b := palindrome[i]
		left := byte(97)
		for left < b {
			newPalindrome := palindrome[:i] + string(rune(left)) + palindrome[i + 1:]
			if !isPalindrome(newPalindrome) {
				return newPalindrome
			}
			left++
		}
	}

	for i := len(palindrome) - 1; i >= 0; i-- {
		b := palindrome[i]
		right := b + 1
		for right <= 122 {
			newPalindrome := palindrome[:i] + string(rune(right)) + palindrome[i + 1:]
			// fmt.Printf("new = %v\n", newPalindrome)
			if !isPalindrome(newPalindrome) {
				return newPalindrome
			}
			right++
		}
	}

	return ""
}

func isPalindrome(s string) bool {
    if len(s) == 1 {
        return true
    }

    i := int(0)
    j := int(len(s) - 1)
    for i <= j {
        left := s[i]
        right := s[j]
        if left != right {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
	// result: "aaccba"
	// palindrome := "abccba"

	// result: ""
	// palindrome := "a"

	// result: "ab"
	// palindrome := "aa"

	// result: "abb"
	// palindrome := "aba"

	// result: "adabccbadz"
	palindrome := "zdabccbadz"

	// result: 
	// palindrome := ""

	result := breakPalindrome(palindrome)
	fmt.Printf("result = %v\n", result)
}

