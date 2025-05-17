package main
import (
	"fmt"
)

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	maxIndex := int(0)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if check(words[i], words[j]) && dp[j] + 1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}

		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	result := make([]string, 0)

	for i := maxIndex; i >= 0; i = prev[i] {
		result = append(result, words[i])
	}

	reverse(result)
	return result
}

func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	diff := int(0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++

			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}

func reverse(arr []string) {
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	// result: ["bab","cab"]
	// words := []string{"bab","dab","cab"}
	// groups := []int{1,2,2}

	// result: ["a","b","c","d"]
	words := []string{"a","b","c","d"}
	groups := []int{1,2,3,4}

	// result: []
	// words := []string{}
	// groups := []int{}

	result := getWordsInLongestSubsequence(words, groups)
	fmt.Printf("result = %v\n", result)
}

