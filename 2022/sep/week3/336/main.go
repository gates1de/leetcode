package main
import (
	"fmt"
)

type TrieNode struct {
	Children          map[byte]*TrieNode
	WordId            int
	PalindromeWordIds []int
}

type Trie struct {
	root *TrieNode
}

func (this *Trie) insert(index int, word string) {
	node := this.root

	for i := 0; i < len(word); i++ {
		c := word[len(word) - 1 - i]
		if isPalindrome(word[:len(word) - i]) {
			node.PalindromeWordIds = append(node.PalindromeWordIds, index)
		}

		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{
				Children: make(map[byte]*TrieNode),
				WordId:   -1,
			}
		}

		node = node.Children[c]
	}

	node.WordId = index
}

func (this *Trie) search(index int, word string) [][]int {
	node := this.root
	result := [][]int{}

	for len(word) > 0 {
		if node.WordId > -1 && isPalindrome(word) {
			result = append(result, []int{index, node.WordId})
		}
		if node.Children[word[0]] == nil {
			return result
		}
		node = node.Children[word[0]]
		word = word[1:]
	}

	if node.WordId > -1 && node.WordId != index {
		result = append(result, []int{index, node.WordId})
	}

	for _, palindromeWordId := range node.PalindromeWordIds {
		result = append(result, []int{index, palindromeWordId})
	}

	return result
}

func isPalindrome(word string) bool {
	n := len(word)
	for i := 0; i < n / 2; i++ {
		if word[i] != word[n - 1 - i] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	trie := Trie{
		&TrieNode{
			Children: make(map[byte]*TrieNode),
			WordId:   -1,
		},
	}

	result := [][]int{}

	for i, w := range words {
		trie.insert(i, w)
	}

	for i, w := range words {
		result = append(result, trie.search(i, w)...)
	}

	return result
}

func main() {
	// result: [[0,1],[1,0],[3,2],[2,4]]
	// words := []string{"abcd","dcba","lls","s","sssll"}

	// result: [[0,1],[1,0]]
	// words := []string{"bat","tab","cat"}

	// result: [[0,1],[1,0]]
	words := []string{"a",""}

	// result: 
	// words := []string{}

	result := palindromePairs(words)
	fmt.Printf("result = %v\n", result)
}

