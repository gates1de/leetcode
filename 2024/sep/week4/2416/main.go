package main
import (
	"fmt"
)

type TrieNode struct {
	Count int
	Next []*TrieNode
}

func (this *TrieNode) Insert(word string) {
	node := this
	for _, char := range word {
		if node.Next[char - 'a'] == nil {
			node.Next[char - 'a'] = &TrieNode{Count: 0, Next: make([]*TrieNode, 26)}
		}
		node.Next[char - 'a'].Count++
		node = node.Next[char - 'a']
	}
}

func (this *TrieNode) CountAll(word string) int {
	node := this
	result := int(0)

	for _, char := range word {
		result += node.Next[char - 'a'].Count
		node = node.Next[char - 'a']
	}
	return result
}

func sumPrefixScores(words []string) []int {
	root := &TrieNode{Count: 0, Next: make([]*TrieNode, 26)}

	n := len(words)
	for _, word := range words {
		root.Insert(word)
	}

	result := make([]int, n)

	for i, _ := range result {
		result[i] = root.CountAll(words[i])
	}

	return result
}

func main() {
	// result: [5,4,3,2]
	// words := []string{"abc","ab","bc","b"}

	// result: [4]
	words := []string{"abcd"}

	// result: []
	// words := []string{}

	result := sumPrefixScores(words)
	fmt.Printf("result = %v\n", result)
}

