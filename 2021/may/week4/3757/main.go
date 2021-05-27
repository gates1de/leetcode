package main
import (
	"fmt"
	"sort"
)

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	alpha := make([]int, len(words))
	for i, w := range words {
		for _, c := range w {
			alpha[i] = alpha[i] | (0x1 << uint(c - 'a'))
		}
	}
	max := 0
	for i := 0; i < len(words); i++ {
		if max >= len(words[i]) * len(words[i]) {
			break
		}
		for j := i + 1; j < len(words); j++ {
			if max >= len(words[i]) * len(words[j]) {
				break
			}
			if alpha[i] & alpha[j] != 0 {
				continue
			}
			max = len(words[i]) * len(words[j])
			break
		}
	}
	return max
}

// Slow & Use more memory
func mySolution(words []string) int {
	dp := map[int]int{}
	for i, word1 := range words {
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if !isCommon(word1, word2) {
				dp[len(word1) * len(word2)]++
			}
		}
	}
	// fmt.Printf("dp = %v\n", dp)
	result := int(0)
	for val, _ := range dp {
		if result < val {
			result = val
		}
	}
	return result
}

func isCommon(word1 string, word2 string) bool {
	m := map[rune]bool{}
	for _, r := range word1 {
		m[r] = true
	}
	for _, r := range word2 {
		if m[r] {
			return true
		}
	}
	return false
}

func main() {
	// result: 16
	// words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}

	// result: 4
	// words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}

	// result: 0
	words := []string{"a", "aa", "aaa", "aaaa"}

	// result: 
	// words := []string{}

	result := maxProduct(words)
	fmt.Printf("result = %v\n", result)
}

