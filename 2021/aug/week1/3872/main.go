package main
import (
	"fmt"
	"math"
	"strings"
)

func minCut(s string) int {
    if len(s) == 1 {
        return 0
    }
    dict := make(map[int]int)
    return helper(s, dict)
}

func helper(s string, dict map[int]int) int {
    if len(s) == 0 {
        return -1
    }

    if v, ok := dict[len(s)]; ok {
        return v
    }

    result := math.MaxInt64
    for i := 1; i <= len(s); i++ {
        if isPalindrome(s[:i]) {
            min := 1 + helper(s[i:], dict)
            if min < result {
                result = min
            }
        }
	}
    dict[len(s)] = result
    return result
}

// Wrong Answer
func ngSolution(s string) int {
	if len(s) == 1 {
		return 0
	}

	result := int(0)
	TOP:
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		lastCharIndex := len(s)
		for lastCharIndex >= 0 {
			lastCharIndex = strings.LastIndex(s[i + 1:lastCharIndex], char)
			// fmt.Printf("char = %v, lastCharIndex = %v\n", char, lastCharIndex)
			if lastCharIndex == -1 {
				result++
				continue TOP
			}

			lastCharIndex += i + 1
			target := s[i:lastCharIndex + 1]
			// fmt.Printf("target = %v\n", target)
			if !isPalindrome(target) {
				continue
			}
			i = lastCharIndex
			break
		}

		result++
	}

	// remove last partition (ex: ["aa" | "b" | <- this]
	result--
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
	// result: 1
	// s := "aab"

	// result: 0
	// s := "a"

	// result: 1
	// s := "ab"

	// result: 1
	// s := "aaabaaaa"

	// result: 4
	s := "abacbccaab"

	// result: 1
	// s := "aaabaa"

	// result: 
	// s := ""

	result := minCut(s)
	fmt.Printf("result = %v\n", result)
}

