package main
import (
	"fmt"
)

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return s
    }

    result := string(s[0])
    max := int(1)
    for i := 0; i < len(s); i++ {
        if max >= len(s) - i {
            break
        }

        for j := i + 1; j <= len(s); j++ {
            target := s[i:j]
            if max >= len(target) {
                continue
            }
            if isPalindrome(target) {
                result = target
                max = len(target)
            }
        }
    }

    return result
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
	// result: "bab"
	// s := "babad"

	// result: "bb"
	s := "cbbd"

	// result: ""
	// s := ""

	result := longestPalindrome(s)
	fmt.Printf("result = %v\n", result)
}

