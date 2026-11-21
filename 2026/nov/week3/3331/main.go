package main

import (
	"fmt"
)

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	if n == 0 {
		return []int{}
	}

	children := make([][]int, n)
	for node := 1; node < n; node++ {
		children[parent[node]] = append(children[parent[node]], node)
	}

	newParent := append([]int(nil), parent...)
	last := [26]int{}
	for i := range last {
		last[i] = -1
	}

	type event struct {
		node  int
		enter bool
		prev  int
	}
	stack := []event{{node: 0, enter: true}}
	order := make([]int, 0, n)

	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		current := stack[lastIndex]
		stack = stack[:lastIndex]
		character := int(s[current.node] - 'a')

		if !current.enter {
			last[character] = current.prev
			continue
		}

		currentAncestor := last[character]
		if currentAncestor != -1 {
			newParent[current.node] = currentAncestor
		}
		last[character] = current.node
		order = append(order, current.node)
		stack = append(stack, event{
			node:  current.node,
			prev:  currentAncestor,
			enter: false,
		})

		for i := len(children[current.node]) - 1; i >= 0; i-- {
			stack = append(stack, event{
				node:  children[current.node][i],
				enter: true,
			})
		}
	}

	result := make([]int, n)
	for node := range result {
		result[node] = 1
	}
	for i := len(order) - 1; i > 0; i-- {
		node := order[i]
		result[newParent[node]] += result[node]
	}

	return result
}

func main() {
	// result: [6,3,1,1,1,1]
	// parent := []int{-1,0,0,1,1,1}
	// s := "abaabc"

	// result: [5,2,1,1,1]
	parent := []int{-1, 0, 4, 0, 1}
	s := "abbba"

	// result: []
	// parent := []int{}
	// s := ""

	result := findSubtreeSizes(parent, s)
	fmt.Printf("result = %v\n", result)
}
