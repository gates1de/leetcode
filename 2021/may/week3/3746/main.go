package main
import (
	"fmt"
	"sort"
	"strings"
)

func longestStrChain(words []string) int {
	if len(words) <= 1 {
		return len(words)
	} else if len(words) == 2 {
		if diffCount(words[0], words[1]) == 1 {
			return 2
		}
		return 1
	}

	dp := map[string]int{}
	sort.Slice(words, func (a, b int) bool { return len(words[a]) < len(words[b]) })
	// fmt.Printf("words = %v\n", words)

	result := int(0)
	for i, word1 := range words {
		// fmt.Printf("topWord = %v\n", word1)
		count := int(1)
		if i + 1 < len(words) {
			helper(word1, word1, count, words[i + 1:], dp)
		}
	}
	// fmt.Printf("dp = %v\n", dp)
	for _, count := range dp {
		if result < count {
			result = count
		}
	}
	return result
}

func helper(topWord string, currentWord string, currentCount int, words []string, dp map[string]int) {
	if len(words) == 0 {
		return
	}

	for i := 0; i < len(words); i++ {
		word2 := words[i]
		diffCount := diffCount(currentWord, word2)
		// fmt.Printf("%v : %v diffCount = %v\n", currentWord, word2, diffCount)

		if diffCount != 1 {
			// fmt.Printf("topWord: %v ~ currentWord: %v = %v\n", topWord, currentWord, currentCount)
			dp[topWord] = max(dp[topWord], currentCount)
			if abs(len(currentWord), len(word2)) >= 2 {
				return
			}
			continue
		}
		dp[topWord] = max(dp[topWord], currentCount + 1)
		if i + 1 < len(words) {
			helper(topWord, word2, currentCount + 1, words[i + 1:], dp)
		}
	}
}

func diffCount(word1 string, word2 string) int {
	if abs(len(word1), len(word2)) != 1 {
		return 0
	}

	// word1 = sortString(word1)
	// word2 = sortString(word2)
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}
	j := int(0)
	result := int(0)
	for _, r1 := range word1 {
		if result >= 2 {
			return result
		}
		if j >= len(word2) {
			result++
			continue
		}
		r2 := rune(word2[j])
		if r1 != r2 {
			result++
			continue
		}
		j++
	}
	return result
}

func sortString(word string) string {
	wordSlice := strings.Split(word, "")
	sort.Strings(wordSlice)
	return strings.Join(wordSlice, "")
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 4
	// words := []string{"a", "b", "ba", "bca", "bda", "bdca"}

	// result: 5
	// words := []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}

	// result: 2
	// words := []string{"a", "b", "ab", "bac"}

	// result: 4
	// words := []string{"bdca", "bda", "ca", "dca", "a"}

	// result: 4
	words := []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"}

	// result: 
	// words := []string{}

	result := longestStrChain(words)
	fmt.Printf("result = %v\n", result)
}

