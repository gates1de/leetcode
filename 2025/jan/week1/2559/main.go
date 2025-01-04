package main
import (
	"fmt"
)

func vowelStrings(words []string, queries [][]int) []int {
	result := make([]int, len(queries))
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	prefixSum := make([]int, len(words))
	sum := int(0)

	for i := 0; i < len(words); i++ {
		currentWord := words[i]

		if vowels[currentWord[0]] && vowels[currentWord[len(currentWord) - 1]] {
			sum++
		}

		prefixSum[i] = sum
	}

	for i := 0; i < len(queries); i++ {
		currentQuery := queries[i]
		result[i] = prefixSum[currentQuery[1]]

		if currentQuery[0] != 0 {
			result[i] -= prefixSum[currentQuery[0] - 1]
		}
	}

	return result
}

func main() {
	// result: [2,3,0]
	// words := []string{"aba","bcb","ece","aa","e"}
	// queries := [][]int{{0,2},{1,4},{1,1}}

	// result: [3,2,1]
	words := []string{"a","e","i"}
	queries := [][]int{{0,2},{0,1},{2,2}}

	// result: []
	// words := []string{}
	// queries := [][]int{}

	result := vowelStrings(words, queries)
	fmt.Printf("result = %v\n", result)
}

