package main
import (
	"fmt"
)

func countSubstrings(s string) int {
    result := 0
    for midCharIndex := 0; midCharIndex < len(s); midCharIndex++ {
        left := midCharIndex
        right := midCharIndex + 1
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
            result++
        }

        left = midCharIndex
        right = midCharIndex
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
            result++
        }
    }
    return result
}

// Slow & Use more memory
func mySolution(s string) int {
	substrings := getSubstrings(s)
	result := int(0)
	for _, substring := range substrings {
		if isPalindrome(substring) {
			result++
		}
	}
	return result
}

func getSubstrings(s string) []string {
	result := []string{}
	for i, rune1 := range s {
		result = append(result, string(rune1))
		if i == len(s) - 1 {
			break
		}
		result = append(result, s[i:])
		str := string(rune1)
		for _, rune2 := range s[i + 1:len(s) - 1] {
			str += string(rune2)
			result = append(result, str)
		}
	}
	return result
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
	s := "abc"
	// s := "aaa"

	result := countSubstrings(s)
	fmt.Printf("result = %v\n", result)
}

