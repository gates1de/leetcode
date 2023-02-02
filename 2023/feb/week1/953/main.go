package main
import (
	"fmt"
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	orderMap := map[byte]int{}
	for i, char := range order {
		orderMap[byte(char)] = i
	}

	sortedWords := make([]string, len(words))
	copy(sortedWords, words)
	sort.Slice(sortedWords, func(a, b int) bool {
		wordA := sortedWords[a]
		wordB := sortedWords[b]

		for i := 0; i < max(len(wordA), len(wordB)); i++ {
			if i >= len(wordA) {
				return true
			} else if i >= len(wordB) {
				return false
			}

			if orderMap[wordA[i]] < orderMap[wordB[i]] {
				return true
			} else if orderMap[wordA[i]] > orderMap[wordB[i]] {
				return false
			}
		}
		return true
	})

	for i, sortedWord := range sortedWords {
		word := words[i]
		if sortedWord != word {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: true
	// words := []string{"hello","leetcode"}
	// order := "hlabcdefgijkmnopqrstuvwxyz"

	// result: false
	// words := []string{"word","world","row"}
	// order := "worldabcefghijkmnpqstuvxyz"

	// result: false
	words := []string{"apple","app"}
	order := "abcdefghijklmnopqrstuvwxyz"

	// result: 
	// words := []string{}
	// order := ""

	result := isAlienSorted(words, order)
	fmt.Printf("result = %v\n", result)
}

