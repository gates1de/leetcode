package main

import (
	"fmt"
)

type UnionFind struct {
	Parents []int
	Ranks 	[]int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := range n {
		parents[i] = i
		ranks[i] = 1
	}

	return &UnionFind{Parents: parents, Ranks: ranks}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parents[x] != x {
		uf.Parents[x] = uf.Find(uf.Parents[x])
	}

	return uf.Parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.Ranks[x] < uf.Ranks[y] {
		x, y = y, x
	}

	uf.Parents[y] = x
	if uf.Ranks[x] == uf.Ranks[y] {
		uf.Ranks[x]++
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := NewUnionFind(n)
	for _, pair := range allowedSwaps {
		uf.Union(pair[0], pair[1])
	}

	sets := make(map[int]map[int]int)
	for i := range n {
		f := uf.Find(i)
		if sets[f] == nil {
			sets[f] = make(map[int]int)
		}
		sets[f][source[i]]++
	}

	result := int(0)
	for i := range n {
		f := uf.Find(i)
		if sets[f][target[i]] > 0 {
			sets[f][target[i]]--
		} else {
			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// source := []int{1,2,3,4}
	// target := []int{2,1,4,5}
	// allowedSwaps := [][]int{{0,1},{2,3}}

	// result: 2
	// source := []int{1,2,3,4}
	// target := []int{1,3,2,4}
	// allowedSwaps := [][]int{}

	// result: 0
	source := []int{5,1,2,4,3}
	target := []int{1,5,4,2,3}
	allowedSwaps := [][]int{{0,4},{4,2},{1,3},{1,4}}

	// result: 
	// source := []int{}
	// target := []int{}
	// allowedSwaps := [][]int{}

	result := minimumHammingDistance(source, target, allowedSwaps)
	fmt.Printf("result = %v\n", result)
}

