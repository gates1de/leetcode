package main
import (
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
    m, n, l := len(s1), len(s2), len(s3)
    if m + n != l {
        return false
    }

    dp := make([]bool, n + 1)
    dp[0] = true

    for i := 1; i <= n; i++ {
        dp[i] = dp[i - 1] && s2[i - 1] == s3[i - 1]
    }

    for i := 1; i <= m; i++ {
        dp[0] = dp[0] && s1[i - 1] == s3[i - 1]
        for j := 1; j <= n; j++ {
            dp[j] = (dp[j] && s1[i - 1] == s3[i + j - 1]) || (dp[j - 1] && s2[j - 1] == s3[i + j - 1])
        }
    }

    return dp[n]
}

func main() {
	// result: true
	// s1 := "aabcc"
	// s2 := "dbbca"
	// s3 := "aadbbcbcac"

	// result: false
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"

	// result: 
	// s1 := ""
	// s2 := ""
	// s3 := ""

	result := isInterleave(s1, s2, s3)
	fmt.Printf("result = %v\n", result)
}

