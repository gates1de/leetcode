package main
import (
	"fmt"
)

type TrieNode struct {
	Children [26]*TrieNode
	Word string
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	trie := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, trie, &result)
		}
    }
    return result
}

func dfs(board [][]byte, i int, j int, trie *TrieNode, result *[]string) {
	char := board[i][j]
    if char == '#' || trie.Children[char - 'a'] == nil {
		return
	}

	trie = trie.Children[char - 'a']
    if trie.Word != "" {
        *result = append(*result, trie.Word)
		trie.Word = ""
    }

    board[i][j] = '#'
    if i > 0 {
		dfs(board, i - 1, j, trie, result)
	}
    if j > 0 {
		dfs(board, i, j - 1, trie, result)
	}
    if i < len(board) - 1 {
		dfs(board, i + 1, j, trie, result)
	}
    if j < len(board[0]) - 1 {
		dfs(board, i, j + 1, trie, result)
	}
    board[i][j] = char
}

func buildTrie(words []string) *TrieNode {
	result := &TrieNode{}
	for _, word := range words {
		node := result
		for _, char := range word {
			i := char - 'a'
			if node.Children[i] == nil {
				node.Children[i] = &TrieNode{}
			}
			node = node.Children[i]
		}
		node.Word = word
	}
	return result
}

func main() {
	// result: ["oath","eat"]
	// board := [][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}}
	// words := []string{"oath","pea","eat","rain"}

	// result: []
	board := [][]byte{{'a','b'},{'c','d'}}
	words := []string{"abcb"}

	// result: 
	// board := [][]byte{}
	// words := []string{}

	result := findWords(board, words)
	fmt.Printf("result = %v\n", result)
}

