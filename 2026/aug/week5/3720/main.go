package main

import (
	"fmt"
)

func lexGreaterPermutation(s string, target string) string {
	counts := [26]int{}
	for i := range s {
		counts[s[i]-'a']++
	}

	result := ""
	for i := range target {
		for character := target[i] - 'a' + 1; character < 26; character++ {
			if counts[character] == 0 {
				continue
			}

			candidate := make([]byte, 0, len(s))
			candidate = append(candidate, target[:i]...)
			candidate = append(candidate, character+'a')
			remaining := counts
			remaining[character]--
			for letter, count := range remaining {
				for ; count > 0; count-- {
					candidate = append(candidate, byte(letter)+'a')
				}
			}

			result = string(candidate)
			break
		}

		letter := target[i] - 'a'
		if counts[letter] == 0 {
			break
		}
		counts[letter]--
	}

	return result
}

func main() {
	// result: "bca"
	// s := "abc"
	// target := "bba"

	// result: "eelt"
	// s := "leet"
	// target := "code"

	// result: ""
	s := "baba"
	target := "bbaa"

	// result: ""
	// s := ""
	// target := ""

	result := lexGreaterPermutation(s, target)
	fmt.Printf("result = %v\n", result)
}
