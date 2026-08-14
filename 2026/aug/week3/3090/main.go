package main

import (
	"fmt"
)

func maximumLengthSubstring(s string) int {
	count := [26]int{}
	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		index := s[right] - 'a'
		count[index]++

		for count[index] > 2 {
			count[s[left]-'a']--
			left++
		}

		if length := right - left + 1; length > result {
			result = length
		}
	}

	return result
}

func main() {
	// result: 4
	// s := "bcbbbcba"

	// result: 2
	s := "aaaa"

	// result:
	// s := ""

	result := maximumLengthSubstring(s)
	fmt.Printf("result = %v\n", result)
}
