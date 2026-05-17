package main

import (
	"fmt"
)

func canReach(arr []int, start int) bool {
	n := len(arr)
	if n == 0 || start < 0 || start >= n {
		return false
	}

	visited := make([]bool, n)
	stack := make([]int, 1, n)
	stack[0] = start

	for len(stack) > 0 {
		idx := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if visited[idx] {
			continue
		}
		if arr[idx] == 0 {
			return true
		}
		visited[idx] = true

		jump := arr[idx]
		next := idx + jump
		if next < n && !visited[next] {
			stack = append(stack, next)
		}

		prev := idx - jump
		if prev >= 0 && !visited[prev] {
			stack = append(stack, prev)
		}
	}

	return false
}

func main() {
	// result: true
	// arr := []int{4, 2, 3, 0, 3, 1, 2}
	// start := int(5)

	// result: true
	// arr := []int{4, 2, 3, 0, 3, 1, 2}
	// start := int(0)

	// result: false
	// arr := []int{3, 0, 2, 1, 2}
	// start := int(2)

	// result: true
	arr := []int{0, 1}
	start := int(1)

	// result:
	// arr := []int{}
	// start := int(0)

	result := canReach(arr, start)
	fmt.Printf("result = %v\n", result)
}
