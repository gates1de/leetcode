package main

import (
	"fmt"
)

func remainingMethods(n int, k int, invocations [][]int) []int {
	graph := make([][]int, n)
	for _, invocation := range invocations {
		graph[invocation[0]] = append(graph[invocation[0]], invocation[1])
	}

	suspicious := make([]bool, n)
	stack := []int{k}
	suspicious[k] = true
	for len(stack) > 0 {
		last := len(stack) - 1
		method := stack[last]
		stack = stack[:last]

		for _, next := range graph[method] {
			if suspicious[next] {
				continue
			}
			suspicious[next] = true
			stack = append(stack, next)
		}
	}

	for _, invocation := range invocations {
		if !suspicious[invocation[0]] && suspicious[invocation[1]] {
			result := make([]int, n)
			for method := range n {
				result[method] = method
			}
			return result
		}
	}

	result := make([]int, 0, n)
	for method := range n {
		if !suspicious[method] {
			result = append(result, method)
		}
	}

	return result
}

func main() {
	// result: [0,1,2,3]
	// n := int(4)
	// k := int(1)
	// invocations := [][]int{{1,2},{0,1},{3,2}}

	// result: [3,4]
	// n := int(5)
	// k := int(0)
	// invocations := [][]int{{1,2},{0,2},{0,1},{3,4}}

	// result: []
	n := int(3)
	k := int(2)
	invocations := [][]int{{1, 2}, {0, 1}, {2, 0}}

	// result: []
	// n := int(0)
	// k := int(0)
	// invocations := [][]int{}

	result := remainingMethods(n, k, invocations)
	fmt.Printf("result = %v\n", result)
}
