package main

import (
	"fmt"
)

type SuffixTrieNode struct {
	Child [26]int32
	Best  int32
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	const noChild = int32(-1)

	nodes := make([]SuffixTrieNode, 1, 1 + len(wordsContainer))
	initSuffixTrieNode(&nodes[0])

	contLens := make([]int, len(wordsContainer))
	for i, word := range wordsContainer {
		contLens[i] = len(word)
		cur := int(0)
		if isBetterIndex(i, int(nodes[cur].Best), contLens) {
			nodes[cur].Best = int32(i)
		}

		for p := len(word) - 1; p >= 0; p-- {
			idx := int(word[p] - 'a')
			next := nodes[cur].Child[idx]
			if next == noChild {
				next = int32(len(nodes))
				nodes[cur].Child[idx] = next
				nodes = append(nodes, SuffixTrieNode{})
				initSuffixTrieNode(&nodes[next])
			}

			cur = int(next)
			if isBetterIndex(i, int(nodes[cur].Best), contLens) {
				nodes[cur].Best = int32(i)
			}
		}
	}

	result := make([]int, len(wordsQuery))
	for i, word := range wordsQuery {
		cur := 0
		best := int(nodes[cur].Best)
		for p := len(word) - 1; p >= 0; p-- {
			idx := int(word[p] - 'a')
			next := nodes[cur].Child[idx]
			if next == noChild {
				break
			}

			cur = int(next)
			best = int(nodes[cur].Best)
		}

		result[i] = best
	}

	return result
}

func isBetterIndex(a, b int, lens []int) bool {
	if a == -1 {
		return false
	}

	if b == -1 {
		return true
	}

	if lens[a] != lens[b] {
		return lens[a] < lens[b]
	}

	return a < b
}

func initSuffixTrieNode(n *SuffixTrieNode) {
	for i := range n.Child {
		n.Child[i] = -1
	}

	n.Best = -1
}

func main() {
	// result: [1,1,1]
	// wordsContainer := []string{"abcd","bcd","xbcd"}
	// wordsQuery := []string{"cd","bcd","xyz"}

	// result: [2,0,2]
	wordsContainer := []string{"abcdefgh","poiuygh","ghghgh"}
	wordsQuery := []string{"gh","acbfgh","acbfegh"}

	// result: []
	// wordsContainer := []string{}
	// wordsQuery := []string{}

	result := stringIndices(wordsContainer, wordsQuery)
	fmt.Printf("result = %v\n", result)
}
