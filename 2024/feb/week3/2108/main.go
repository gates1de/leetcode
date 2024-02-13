package main
import (
	"fmt"
)

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}

func isPalindrome(s string) bool {
    if len(s) == 0 {
        return false
    } else if len(s) == 1 {
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
	// result: "ada"
	// words := []string{"abc","car","ada","racecar","cool"}

	// result: "racecar"
	// words := []string{"notapalindrome","racecar"}

	// result: ""
	words := []string{"def","ghi"}

	// result: ""
	// words := []string{}

	result := firstPalindrome(words)
	fmt.Printf("result = %v\n", result)
}

