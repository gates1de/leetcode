package main

import (
	"fmt"
	"sort"
	"strings"
)

type Trie struct {
	Serial   string
	Children map[string]*Trie
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := &Trie{Children: make(map[string]*Trie)}

	for _, path := range paths {
		current := root

		for _, node := range path {
			if _, exists := current.Children[node]; !exists {
				current.Children[node] = &Trie{Children: make(map[string]*Trie)}
			}

			current = current.Children[node]
		}
	}

	freq := make(map[string]int)
	construct(root, freq)

	result := make([][]string, 0)
	path := make([]string, 0)
	operate(root, freq, path, &result)

	return result
}

func construct(node *Trie, freq map[string]int) *Trie {
	if len(node.Children) == 0 {
		return node
	}

	v := make([]string, 0, len(node.Children))
	for folder, child := range node.Children {
		child = construct(child, freq)
		v = append(v, fmt.Sprintf("%s(%s)", folder, child.Serial))
	}

	sort.Strings(v)
	node.Serial = strings.Join(v, "")
	freq[node.Serial]++

	return node
}

func operate(node *Trie, freq map[string]int, path []string, result *[][]string) {
	if freq[node.Serial] > 1 {
		return
	}

	if len(path) > 0 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}

	for folder, child := range node.Children {
		path = append(path, folder)
		operate(child, freq, path, result)
		path = path[:len(path)-1]
	}
}

func main() {
	// result: [["d"],["d","a"]]
	// paths := [][]string{{}}"a"},{"c"},{}"d"},{}"a","b"},{}"c","b"},{}"d","a"}}

	// result: [["c"],["c","b"],["a"],["a","b"]]
	// paths := [][]string{{}}"a"},{}"c"},{}"a","b"},{}"c","b"},{}"a","b","x"},{}"a","b","x","y"},{}"w"},{}"w","y"}}

	// result: [["c"],["c","d"],["a"],["a","b"]]
	paths := [][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}}

	// result: []
	// paths := [][]string{}

	result := deleteDuplicateFolder(paths)
	fmt.Printf("result = %v\n", result)
}
