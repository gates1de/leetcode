package main

import (
	"fmt"
)

func maxPalindromes(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	palindrome := make([]bool, n)

	for right := range s {
		for left := range palindrome[:right+1] {
			palindrome[left] = s[left] == s[right] &&
				(right-left < 2 || palindrome[left+1])
		}

		dp[right+1] = dp[right]
		if right+1 >= k {
			for left := range palindrome[:right-k+2] {
				if palindrome[left] && dp[left]+1 > dp[right+1] {
					dp[right+1] = dp[left] + 1
				}
			}
		}
	}

	return dp[n]
}

func main() {
	// result: 2
	// s := "abaccdbbd"
	// k := int(3)

	// result: 0
	s := "adbcda"
	k := int(2)

	// result: 0
	// s := ""
	// k := int(0)

	result := maxPalindromes(s, k)
	fmt.Printf("result = %v\n", result)
}
