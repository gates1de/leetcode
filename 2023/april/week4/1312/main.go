package main
import (
	"fmt"
)

func minInsertions(s string) int {
	n := len(s)
	bytes := []byte(s)
	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	reversedS := string(bytes)

	memo := make([][]int, n + 1)
	for i, _ := range memo {
		memo[i] = make([]int, n + 1)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}

	return n - lcs(s, reversedS, n, n, memo)
}

func lcs(s1 string, s2 string, m int, n int, memo [][]int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if s1[m - 1] == s2[n - 1] {
		memo[m][n] = 1 + lcs(s1, s2, m - 1, n - 1, memo)
		return memo[m][n]
	}

	memo[m][n] = max(lcs(s1, s2, m - 1, n, memo), lcs(s1, s2, m, n - 1, memo))
	return memo[m][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// s := "zzazz"

	// result: 2
	// s := "mbadm"

	// result: 5
	s := "leetcode"

	// result: 
	// s := ""

	result := minInsertions(s)
	fmt.Printf("result = %v\n", result)
}

