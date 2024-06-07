package main
import (
	"fmt"
	"strings"
)

type TrieNode struct {
	IsEnd bool
	Children [26]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func (this *Trie) Insert(word string) {
	current := this.Root
	for _, char := range word {
		if current.Children[char - 'a'] == nil {
			current.Children[char - 'a'] = &TrieNode{IsEnd: false, Children: [26]*TrieNode{}}
		}
		current = current.Children[char - 'a']
	}
	current.IsEnd = true
}

func (this *Trie) ShortestRoot(word string) string {
	current := this.Root
	for i, char := range word {
		if current.Children[char - 'a'] == nil {
			return word
		}

		current = current.Children[char - 'a']
		if current.IsEnd {
			return word[:i + 1]
		}
	}

	return word
}

func replaceWords(dictionary []string, sentence string) string {
	wordArray := strings.Split(sentence, " ")
	trie := Trie{Root: &TrieNode{IsEnd: false, Children: [26]*TrieNode{}}}
	for _, word := range dictionary {
		trie.Insert(word)
	}

	for i, _ := range wordArray {
		wordArray[i] = trie.ShortestRoot(wordArray[i])
	}

	return strings.Join(wordArray, " ")
}

func main() {
	// result: "the cat was rat by the bat"
	// dictionary := []string{"cat","bat","rat"}
	// sentence := "the cattle was rattled by the battery"

	// result: "a a b c"
	dictionary := []string{"a","b","c"}
	sentence := "aadsfasf absbs bbab cadsfafs"

	// result: ""
	// dictionary := []string{}
	// sentence := ""

	result := replaceWords(dictionary, sentence)
	fmt.Printf("result = %v\n", result)
}

