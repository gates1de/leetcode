package main
import (
	"fmt"
)

func removePalindromeSub(s string) int {
	if isPalindrome(s) {
		return 1
	}
	return 2
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
	// result: 1
	// s := "ababa"

	// result: 2
	// s := "abb"

	// result: 2
	// s := "baabb"

	// result: 2
	// s := "baabbbaabbbaabbbaabbbaabbbaabb"

	// result: 2
	// s := "abbbaa"

	// result: 2
	// s := "ababbabbababaa"

	// result: 2
	s := "ababababababaabbabbbababababaaaabababbbbbbabbabababababbaaa"

	// result: 
	// s := ""

	result := removePalindromeSub(s)
	fmt.Printf("result = %v\n", result)
}

