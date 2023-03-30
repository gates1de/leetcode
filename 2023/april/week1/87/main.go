package main
import (
	"fmt"
)

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]bool, n + 1)
	for i, _ := range dp {
		dp[i] = make([][]bool, n)
		for j, _ := range dp[i] {
			dp[i][j] = make([]bool, n)
		}
	}

	for i, _ := range dp[1] {
		for j, _ := range dp[1][i] {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n + 1 - length; i++ {
			for j := 0; j < n + 1 - length; j++ {
				for newLength := 1; newLength < length; newLength++ {
					dp1 := dp[newLength][i]
					dp2 := dp[length - newLength][i + newLength]
					dp[length][i][j] = dp[length][i][j] || (dp1[j] && dp2[j + newLength])
					dp[length][i][j] = dp[length][i][j] || (dp1[j + length - newLength] && dp2[j])
				}
			}
		}
	}
	return dp[n][0][0]
}

func main() {
	// result: true
	// s1 := "great"
	// s2 := "rgeat"

	// result: false
	// s1 := "abcde"
	// s2 := "caebd"

	// result: true
	s1 := "a"
	s2 := "a"

	// result: 
	// s1 := ""
	// s2 := ""

	result := isScramble(s1, s2)
	fmt.Printf("result = %v\n", result)
}

