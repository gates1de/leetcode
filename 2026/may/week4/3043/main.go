package main

import (
	"fmt"
)

type TrieNode struct {
	Children [10]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func highestDivisor(n int) int {
	if n == 0 {
		return 1
	}
	div := 1
	for n >= 10 {
		n /= 10
		div *= 10
	}
	return div
}

func (t *Trie) Insert(num int) {
	node := t.Root
	div := highestDivisor(num)

	for div > 0 {
		digit := (num / div) % 10
		if node.Children[digit] == nil {
			node.Children[digit] = &TrieNode{}
		}
		node = node.Children[digit]
		div /= 10
	}
}

func (t *Trie) FindLongestPrefix(num int) int {
	node := t.Root
	div := highestDivisor(num)
	length := 0

	for div > 0 {
		digit := (num / div) % 10
		next := node.Children[digit]
		if next == nil {
			break
		}
		length++
		node = next
		div /= 10
	}

	return length
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	trie := &Trie{Root: &TrieNode{}}
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

func main() {
	// result: 3
	// arr1 := []int{1,10,100}
	// arr2 := []int{1000}

	// result: 0
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 4, 4}

	// result:
	// arr1 := []int{}
	// arr2 := []int{}

	result := longestCommonPrefix(arr1, arr2)
	fmt.Printf("result = %v\n", result)
}
