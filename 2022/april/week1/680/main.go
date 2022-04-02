package main
import (
	"fmt"
)

func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	return helper(false, 0, len(s) - 1, s)
}

func helper(isAlreadySkipped bool, i int, j int, s string) bool {
	if i >= j {
		return true
	}

	head := s[i]
	tail := s[j]

	if head == tail {
		return helper(isAlreadySkipped, i + 1, j - 1, s)
	}

	if isAlreadySkipped {
		return false
	}

	if helper(true, i + 1, j, s) {
		return true
	}

	return helper(true, i, j - 1, s)
}

func main() {
	// result: true
	// s := "aba"

	// result: true
	// s := "abca"

	// result: false
	// s := "abc"

	// result: true
	// s := "a"

	// result: true
	// s := "zzz"

	// result: true
	// s := "abcca"

	// result: false
	s := "abccaa"

	// result: 
	// s := ""

	result := validPalindrome(s)
	fmt.Printf("result = %v\n", result)
}

