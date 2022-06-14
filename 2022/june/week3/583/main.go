package main
import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
    dp := make([]int, len(word2))
    for i := 0; i < len(word1); i++ {
		prev := int(0)
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

func main() {
	// result: 2
	// word1 := "sea"
	// word2 := "eat"

	// result: 4
	// word1 := "leetcode"
	// word2 := "etco"

	// result: 6
	// word1 := "abc"
	// word2 := "def"

	// result: 4
	word1 := "sea"
	word2 := "etc"

	// result: 
	// word1 := ""
	// word2 := ""

	result := minDistance(word1, word2)
	fmt.Printf("result = %v\n", result)
}

