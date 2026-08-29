package main

import (
	"fmt"
)

func lexPalindromicPermutation(s string, target string) string {
	counts := [26]int{}
	for i := range s {
		counts[s[i] - 'a']++
	}

	middle := byte(0)
	for i, count := range counts {
		if count % 2 == 1 {
			if middle != 0 {
				return ""
			}

			middle = byte(i) + 'a'
		}

		counts[i] /= 2
	}

	buildPalindrome := func(left []byte) string {
		palindrome := make([]byte, 0, len(s))
		palindrome = append(palindrome, left...)
		if middle != 0 {
			palindrome = append(palindrome, middle)
		}

		for i := len(left) - 1; i >= 0; i-- {
			palindrome = append(palindrome, left[i])
		}

		return string(palindrome)
	}

	halfLength := len(s) / 2
	matchingHalf := []byte(target[:halfLength])
	remaining := counts
	canMatch := true
	for _, character := range matchingHalf {
		letter := character - 'a'
		if remaining[letter] == 0 {
			canMatch = false
			break
		}
		remaining[letter]--
	}

	if canMatch && buildPalindrome(matchingHalf) > target {
		return buildPalindrome(matchingHalf)
	}

	result := ""
	for i := range halfLength {
		for character := target[i] - 'a' + 1; character < 26; character++ {
			if counts[character] == 0 {
				continue
			}

			left := make([]byte, 0, halfLength)
			left = append(left, target[:i]...)
			left = append(left, character+'a')
			remaining := counts
			remaining[character]--
			for letter, count := range remaining {
				for ; count > 0; count-- {
					left = append(left, byte(letter) + 'a')
				}
			}

			result = buildPalindrome(left)
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
	// result: "baab"
	// s := "baba"
	// target := "abba"

	// result: ""
	// s := "baba"
	// target := "bbaa"

	// result: ""
	// s := "abc"
	// target := "abb"

	// result: "aca"
	s := "aac"
	target := "abb"

	// result: ""
	// s := ""
	// target := ""

	result := lexPalindromicPermutation(s, target)
	fmt.Printf("result = %v\n", result)
}