package main
import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) >= len(text2) {
		return helper(text2, text1)
	}
	return helper(text1, text2)
}

func helper(s1 string, s2 string) int {
	dp := make([]int, len(s1) + 1)

	for i := 1; i <= len(s2); i++ {
		prev := int(0)

		for j := 1; j <= len(s1); j++ {
			temp := dp[j]

			if s1[j - 1] == s2[i - 1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j - 1], dp[j])
			}

			prev = temp
		}
	}

	return dp[len(s1)]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// text1 := "abcde"
	// text2 := "ace"

	// result: 3
	// text1 := "abc"
	// text2 := "abc"

	// result: 0
	text1 := "abc"
	text2 := "def"

	// result: 
	// text1 := ""
	// text2 := ""

	result := longestCommonSubsequence(text1, text2)
	fmt.Printf("result = %v\n", result)
}

