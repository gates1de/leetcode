package main
import (
	"fmt"
	"strconv"
)

type TrieNode struct {
	Children []*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func (this *Trie) Insert(num int) {
	root := this.Root
	numStr := strconv.Itoa(num)

	for _, digit := range numStr {
		index := digit - '0'

		if root.Children[index] == nil {
			root.Children[index] = &TrieNode{Children: make([]*TrieNode, 10)}
		}
		root = root.Children[index]
	}
}

func (this *Trie) FindLongestPrefix(num int) int {
	root := this.Root
	numStr := strconv.Itoa(num)
	result := int(0)

	for _, digit := range numStr {
		index := digit - '0'

		if root.Children[index] != nil {
			result++
			root = root.Children[index]
		} else {
			break
		}
	}

	return result
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	trie := &Trie{Root: &TrieNode{Children: make([]*TrieNode, 10)}}
	for _, num := range arr1 {
		trie.Insert(num)
	}

	result := int(0)
	for _, num := range arr2 {
		length := trie.FindLongestPrefix(num)
		result = max(result, length)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// arr1 := []int{1,10,100}
	// arr2 := []int{1000}

	// result: 0
	arr1 := []int{1,2,3}
	arr2 := []int{4,4,4}

	// result: 
	// arr1 := []int{}
	// arr2 := []int{}

	result := longestCommonPrefix(arr1, arr2)
	fmt.Printf("result = %v\n", result)
}

