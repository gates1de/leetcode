package main
import (
	"fmt"
)

func removeOccurrences(s string, part string) string {
	n := len(s)
	// kmp = Knuth-Morris-Pratt Algorithm
	kmpLongestPrefixSuffix := computeLongestPrefixSuffix(part)

	stack := make([]byte, 0, n)
	patternIndexes := make([]int, n + 1)

	patternIndex := int(0)
	for i := 0; i < len(s); i++ {
		char := s[i]
		stack = append(stack, char)

		if char == part[patternIndex] {
			patternIndex++
			patternIndexes[len(stack)] = patternIndex

			if patternIndex == len(part) {
				remainingLength := len(part)
				for remainingLength != 0 {
					stack = stack[:len(stack) - 1]
					remainingLength--
				}

				patternIndex = int(0)
				if len(stack) > 0 {
					patternIndex = patternIndexes[len(stack)]
				}
			}
		} else {
			if patternIndex != 0 {
				i--
				patternIndex = kmpLongestPrefixSuffix[patternIndex - 1]
				stack = stack[:len(stack) - 1]
			} else {
				patternIndexes[len(stack)] = 0
			}
		}
	}

	result := make([]byte, 0, len(stack))
	for len(stack) > 0 {
		result = append(result, stack[len(stack) - 1])
		stack = stack[:len(stack) - 1]
	}

	reverse(result)
	return string(result)
}

func computeLongestPrefixSuffix(pattern string) []int {
	result := make([]int, len(pattern))
	i := int(1)
	prefixLen := int(0)

	for i < len(pattern) {
		if pattern[i] == pattern[prefixLen] {
			prefixLen++
			result[i] = prefixLen
			i++
		} else if prefixLen != 0 {
			prefixLen = result[prefixLen - 1]
		} else {
			result[i] = 0
			i++
		}
	}

	return result
}

func reverse(arr []byte) {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - i - 1] = arr[len(arr) - i - 1], arr[i]
	}
}

func main() {
	// result: "dab"
	// s := "daabcbaabcbc"
	// part := "abc"

	// result: "ab"
	s := "axxxxyyyyb"
	part := "xy"

	// result: ""
	// s := ""
	// part := ""

	result := removeOccurrences(s, part)
	fmt.Printf("result = %v\n", result)
}

