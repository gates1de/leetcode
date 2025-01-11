package main
import (
	"fmt"
)

type TrieNode struct {
	Freq int
	Children map[rune]*TrieNode
}

func (this *TrieNode) Insert(word string) {
	for _, c := range word {
		t := TrieNode{}
		if this.Children == nil {
			this.Children = make(map[rune]*TrieNode)
		}

		if this.Children[c] == nil {
			this.Children[c] = &t
		}

		this = this.Children[c]
		this.Freq++
	}
}

func (this *TrieNode) IsSubstring(word string) bool {
	for _, c := range word {
		this = this.Children[c]
	}

	return this.Freq > 1
}

func stringMatching(words []string) []string {
	n := len(words)
	result := make([]string, 0, n)
	root := &TrieNode{}

	for _, word := range words {
		for start := 0; start < len(word); start++ {
			root.Insert(string(word[start:]))
		}
	}

	for _, word := range words {
		if root.IsSubstring(word) {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	// result: ["as","hero"]
	// words := []string{"mass","as","hero","superhero"}

	// result: ["et","code"]
	// words := []string{"leetcode","et","code"}

	// result: []
	words := []string{"blue","green","bu"}

	// result: []
	// words := []string{}

	result := stringMatching(words)
	fmt.Printf("result = %v\n", result)
}

