package main

import (
	"fmt"
)

func shortestBeautifulSubstring(s string, k int) string {
	left, ones := 0, 0
	result := ""

	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			ones++
		}

		for ones > k {
			if s[left] == '1' {
				ones--
			}
			left++
		}

		if ones != k {
			continue
		}

		for s[left] == '0' {
			left++
		}
		candidate := s[left : right+1]
		if result == "" || len(candidate) < len(result) || len(candidate) == len(result) && candidate < result {
			result = candidate
		}
	}

	return result
}

func main() {
	// result: "11001"
	// s := "100011001"
	// k := int(3)

	// result: "11"
	// s := "1011"
	// k := int(2)

	// result: ""
	s := "000"
	k := int(1)

	// result: ""
	// s := ""
	// k := int(0)

	result := shortestBeautifulSubstring(s, k)
	fmt.Printf("result = %v\n", result)
}
