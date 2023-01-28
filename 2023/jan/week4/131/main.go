package main
import (
	"fmt"
)

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	result := make([][]string, 0, 10000)
	dfs(0, make([]string, 0, 20), s, dp, &result)
	return result
}

func dfs(start int, currentList []string, s string, dp [][]bool, result *[][]string) {
	if start >= len(s) {
		*result = append(*result, currentList)
		return
	}

	for end := start; end < len(s); end++ {
		if s[start] == s[end] && (end - start <= 2 || dp[start + 1][end - 1]) {
			dp[start][end] = true
			newCurrentList := make([]string, len(currentList) + 1)
			copy(newCurrentList, currentList)
			newCurrentList[len(newCurrentList) - 1] = string(s[start:end + 1])
			dfs(end + 1, newCurrentList, s, dp, result)
		}
	}
}

func main() {
	// result: [["a","a","b"],["aa","b"]]
	// s := "aab"

	// result: [["a"]]
	s := "a"

	// result: 
	// s := ""

	result := partition(s)
	fmt.Printf("result = %v\n", result)
}

