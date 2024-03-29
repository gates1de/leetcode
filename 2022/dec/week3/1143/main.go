package main
import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
    l1 := len(text1)
    l2 := len(text2)
    dp := make([][]int, l1 + 1)
    for i := range dp {
        dp[i] = make([]int, l2 + 1)
    }
    for i := 1; i <= l1; i++ {
        for j := 1; j <= l2; j++ {
            if text1[i - 1] == text2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1
            } else {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
            }
        }
    }
    return dp[l1][l2]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
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

