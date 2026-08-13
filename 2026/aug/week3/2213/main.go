package main

import (
	"fmt"
)

type node struct {
	left, right                  byte
	prefix, suffix, best, length int
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}

	merge := func(a, b node) node {
		if a.length == 0 {
			return b
		}
		if b.length == 0 {
			return a
		}

		result := node{
			left:   a.left,
			right:  b.right,
			prefix: a.prefix,
			suffix: b.suffix,
			best:   a.best,
			length: a.length + b.length,
		}

		if a.right == b.left {
			if a.prefix == a.length {
				result.prefix += b.prefix
			}
			if b.suffix == b.length {
				result.suffix += a.suffix
			}
			if joined := a.suffix + b.prefix; joined > result.best {
				result.best = joined
			}
		}

		if b.best > result.best {
			result.best = b.best
		}

		return result
	}

	size := int(1)
	for size < n {
		size <<= 1
	}

	tree := make([]node, size*2)
	for i := range n {
		c := s[i]
		tree[size+i] = node{left: c, right: c, prefix: 1, suffix: 1, best: 1, length: 1}
	}

	for i := size - 1; i > 0; i-- {
		tree[i] = merge(tree[i<<1], tree[i<<1|1])
	}

	result := make([]int, len(queryIndices))
	for i, index := range queryIndices {
		c := queryCharacters[i]
		pos := size + index
		tree[pos] = node{left: c, right: c, prefix: 1, suffix: 1, best: 1, length: 1}
		for pos >>= 1; pos > 0; pos >>= 1 {
			tree[pos] = merge(tree[pos<<1], tree[pos<<1|1])
		}
		result[i] = tree[1].best
	}

	return result
}

func main() {
	// result: [3,3,4]
	// s := "babacc"
	// queryCharacters := "bcb"
	// queryIndices := []int{1, 3, 3}

	// result: [2,3]
	s := "abyzz"
	queryCharacters := "aa"
	queryIndices := []int{2, 1}

	// result: []
	// s := ""
	// queryCharacters := ""
	// queryIndices := []int{}

	result := longestRepeating(s, queryCharacters, queryIndices)
	fmt.Printf("result = %v\n", result)
}
