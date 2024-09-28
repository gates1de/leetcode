package main

import (
	"fmt"
)

type Trie struct {
	Children map[byte]*Trie
	IsWord   bool
}

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	root := &Trie{Children: make(map[byte]*Trie), IsWord: false}

	for _, word := range dictionary {
		node := root

		for _, c := range word {
			b := byte(c)
			if node.Children[b] == nil {
				node.Children[b] = &Trie{Children: make(map[byte]*Trie), IsWord: false}
			}

			node = node.Children[b]
		}

		node.IsWord = true
	}

	memo := make([]int, n + 1)
	for i, _ := range memo {
		memo[i] = -1
	}

	return dp(0, n, s, root, &memo)
}

func dp(start int, n int, s string, root *Trie, memo *[]int) int {
	if start == n {
		return 0
	}

	if (*memo)[start] >= 0 {
		return (*memo)[start]
	}

	node := root
	result := dp(start + 1, n, s, root, memo) + 1
	for end := start; end < n; end++ {
		c := s[end]
		if node.Children[c] == nil {
			break
		}

		node = node.Children[c]
		if node.IsWord {
			result = min(result, dp(end + 1, n, s, root, memo))
		}
	}

	(*memo)[start] = result
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// s := "leetscode"
	// dictionary := []string{"leet","code","leetcode"}

	// result: 3
	s := "sayhelloworld"
	dictionary := []string{"hello", "world"}

	// result:
	// s := ""
	// dictionary := []string{}

	result := minExtraChar(s, dictionary)
	fmt.Printf("result = %v\n", result)
}
