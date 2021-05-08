package main
import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2))
	for i := 0; i < len(word1); i++ {
		var prev int
		for j := 0; j < len(word2); j++ {
			temp := dp[j]
			if word1[i] == word2[j] {
				dp[j] = prev + 1
			} else if j > 0 && dp[j - 1] > dp[j] {
				dp[j] = dp[j - 1]
			}
			prev = temp
		}
	}
	return len(word1) + len(word2) - 2 * dp[len(word2) - 1]
}

// Wrong Answer
func ngSolution(word1 string, word2 string) int {
	shortWord := word1
	longWord := word2
	if len(word2) < len(word1) {
		shortWord = word2
		longWord = word1
	}
	maxSameCount := int(0)
	result := int(0)
	blank := generateBlank(len(shortWord) - 1)
	longWordWithBlank := blank + longWord + blank
	// fmt.Printf("shortWord = %v, longWord = %v\n", shortWord, longWord)
	for i := 0; i < len(longWordWithBlank); i++ {
		lRunes := longWordWithBlank[i:min(len(longWordWithBlank), i + len(shortWord))]
		sameCount := sameCounter([]rune(lRunes), []rune(shortWord))
		if maxSameCount < sameCount {
			maxSameCount = sameCount
			result = (len(longWord) - maxSameCount) + (len(shortWord) - maxSameCount)
		}
		// fmt.Printf("lRunes = %v, shortWord = %v, sameCount = %v, result = %v\n", string(lRunes), string([]rune(shortWord)), sameCount, result)
	}
	if maxSameCount == 0 {
		return len(longWord) * 2
	}
	return result
}

func generateBlank(num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += " "
	}
	return result
}

func sameCounter(r1Arr, r2Arr []rune) int {
	result := int(0)
	for i := 0; i < len(r1Arr); i++ {
		if i >= len(r2Arr) {
			break
		}
		r1 := r1Arr[i]
		for j := i; j < len(r2Arr); j++ {
			r2 := r2Arr[j]
			if r1 == r2 {
				result++
				break
			}
		}
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// word1 := "sea"
	// word2 := "eat"

	// result: 4 
	// word1 := "leetcode"
	// word2 := "etco"

	// result: 4
	// word1 := "sea"
	// word2 := "ate"

	// result: 2
	// word1 := "a"
	// word2 := "b"

	// result: 3
	word1 := "park"
	word2 := "spake"

	// result: 
	// word1 := ""
	// word2 := ""

	result := minDistance(word1, word2)
	fmt.Printf("result = %v\n", result)
}

