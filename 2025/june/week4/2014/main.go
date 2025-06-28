package main

import (
	"fmt"
)

func longestSubsequenceRepeatedK(s string, k int) string {
	alphabets := make([]int, 26)
	for _, ch := range s {
		alphabets[ch-'a']++
	}

	candidates := make([]byte, 0, 25)
	for i := 25; i >= 0; i-- {
		if alphabets[i] >= k {
			candidates = append(candidates, byte('a'+i))
		}
	}

	queue := []string{}
	for _, char := range candidates {
		queue = append(queue, string(char))
	}

	result := ""
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current) > len(result) {
			result = current
		}

		for _, c := range candidates {
			next := current + string(c)
			if isKRepeated(s, next, k) {
				queue = append(queue, next)
			}
		}
	}

	return result
}

func isKRepeated(s, pattern string, k int) bool {
	i := int(0)
	matched := int(0)

	for _, char := range s {
		if char != rune(pattern[i]) {
			continue
		}

		i++
		if i == len(pattern) {
			i = 0
			matched++

			if matched == k {
				return true
			}
		}
	}

	return false
}

func main() {
	// result: "let"
	// s := "letsleetcode"
	// k := int(2)

	// result: "b"
	// s := "bb"
	// k := int(2)

	// result: ""
	s := "ab"
	k := int(2)

	// result: ""
	// s := ""
	// k := int(0)

	result := longestSubsequenceRepeatedK(s, k)
	fmt.Printf("result = %v\n", result)
}
