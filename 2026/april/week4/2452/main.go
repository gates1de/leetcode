package main

import (
	"fmt"
)

type TrieNode struct {
	Children [26]*TrieNode
	IsEnd			bool
}

func (this *TrieNode) Insert(word string) {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &TrieNode{}
		}
		node = node.Children[idx]
	}
	node.IsEnd = true
}

func (this *TrieNode) DFS(word string, i int, count int) bool {
	if count > 2 || this == nil {
		return false
	}

	if i == len(word) {
		return this.IsEnd
	}

	idx := word[i] - 'a'

	if this.Children[idx] != nil && this.Children[idx].DFS(word, i + 1, count) {
		return true
	}

	if count < 2 {
		for char := range 26 {
			if byte(char) == idx {
				continue
			}

			if this.Children[char] != nil && this.Children[char].DFS(word, i + 1, count + 1) {
				return true
			}
		}
	}

	return false
}

func twoEditWords(queries []string, dictionary []string) []string {
	root := &TrieNode{}
	for _, word := range dictionary {
		root.Insert(word)
	}

	result := make([]string, 0)
	for _, q := range queries {
		if root.DFS(q, 0, 0) {
			result = append(result, q)
		}
	}

	return result
}

func main() {
	// result: ["word","note","ants","wood"]
	// queries := []string{"word","note","ants","wood"}
	// dictionary := []string{"wood","joke","moat"}

	// result: []
	queries := []string{"yes"}
	dictionary := []string{"not"}

	// result: []
	// queries := []string{}
	// dictionary := []string{}

	result := twoEditWords(queries, dictionary)
	fmt.Printf("result = %v\n", result)
}

