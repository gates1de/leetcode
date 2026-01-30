package main

import (
	"fmt"
	"math"
)

type Trie struct {
	Children [26]*Trie
	ID			 int
}

func newTrie() *Trie {
	return &Trie{ID: -1}
}

func (node *Trie) Add(word string, index *int) int {
	for _, ch := range word {
		i := ch - 'a'
		if node.Children[i] == nil {
			node.Children[i] = newTrie()
		}

		node = node.Children[i]
	}

	if node.ID == -1 {
		*index++
		node.ID = *index
	}

	return node.ID
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	m := len(original)
	root := newTrie()
	p := int(-1)
	nodeCount := m * 2
	counter := make([][]int, nodeCount)

	for i := range counter {
		counter[i] = make([]int, nodeCount)
		for j := range counter[i] {
			counter[i][j] = math.MaxInt32 / 2
		}
		counter[i][i] = 0
	}

	for i := 0; i < m; i++ {
		x := root.Add(original[i], &p)
		y := root.Add(changed[i], &p)
		counter[x][y] = min(counter[x][y], cost[i])
	}

	size := p + 1
	for k := range size {
		for i := range size {
			for j := range size {
				counter[i][j] = min(counter[i][j], counter[i][k] + counter[k][j])
			}
		}
	}

	f := make([]int64, n)
	for i := range f {
		f[i] = -1
	}

	for j := 0; j < n; j++ {
		if j > 0 && f[j - 1] == -1 {
			continue
		}

		base := int64(0)
		if j == 0 {
			base = 0
		} else {
			base = f[j-1]
		}

		if source[j] == target[j] {
			update(&f[j], base)
		}

		u := root
		v := root
		for i := j; i < n; i++ {
			u = u.Children[source[i] - 'a']
			v = v.Children[target[i] - 'a']
			if u == nil || v == nil {
				break
			}

			if u.ID != -1 && v.ID != -1 && counter[u.ID][v.ID] != math.MaxInt32 / 2 {
				newVal := base + int64(counter[u.ID][v.ID])
				update(&f[i], newVal)
			}
		}
	}

	return f[n - 1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func update(x *int64, y int64) {
	if *x == -1 || y < *x {
		*x = y
	}
}

func main() {
	// result: 28
	// source := "abcd"
	// target := "acbe"
	// original := []string{"a","b","c","c","e","d"}
	// changed := []string{"b","c","b","e","b","e"}
	// cost := []int{2,5,5,1,2,20}

	// result: 9
	// source := "abcdefgh"
	// target := "acdeeghh"
	// original := []string{"bcd","fgh","thh"}
	// changed := []string{"cde","thh","ghh"}
	// cost := []int{1,3,5}

	// result: -1
	source := "abcdefgh"
	target := "addddddd"
	original := []string{"bcd","defgh"}
	changed := []string{"ddd","ddddd"}
	cost := []int{100,1578}

	// result: 
	// source := ""
	// target := ""
	// original := []string{}
	// changed := []string{}
	// cost := []int{}

	result := minimumCost(source, target, original, changed, cost)
	fmt.Printf("result = %v\n", result)
}

