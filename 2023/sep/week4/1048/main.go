package main
import (
	"fmt"
	"sort"
)

func longestStrChain(words []string) int {
    n := len(words)
    dp := make([]int, n)
    result := int(1)
    sort.SliceStable(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })

    lenPositions := make(map[int]int)
    for i := range words {
        lenPositions[len(words[i])] = i
    }

    for i := 0; i < n; i++ {
        currMax := int(1)
        if pos, ok := lenPositions[len(words[i]) - 1]; ok {
            for j := pos; j >= 0 && len(words[i]) - len(words[j]) == 1; j-- {
                if isValid(words[j], words[i]) {
                    currMax = max(currMax, dp[j] + 1)
                }
            }
        }
        dp[i] = currMax
        result = max(result, currMax)
    }

    return result
}

func isValid(s1, s2 string) bool {
    i := 0
    for i < len(s1) {
        if s1[i] != s2[i] {
            break
        }
        i++
    }
    return s1 == s2[0:i]+s2[i+1:]
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 4
	// words := []string{"a","b","ba","bca","bda","bdca"}

	// result: 5
	// words := []string{"xbc","pcxbcf","xb","cxbc","pcxbc"}

	// result: 1
	words := []string{"abcd","dbqca"}

	// result: 
	// words := []string{}

	result := longestStrChain(words)
	fmt.Printf("result = %v\n", result)
}

