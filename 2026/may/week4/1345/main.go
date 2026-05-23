package main

import (
	"fmt"
)

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	indicesByValue := make(map[int][]int, n)
	for i, v := range arr {
		indicesByValue[v] = append(indicesByValue[v], i)
	}

	visited := make([]bool, n)
	visited[0] = true

	queue := make([]int, 1, n)
	queue[0] = 0
	head := 0
	steps := 0

	for head < len(queue) {
		levelEnd := len(queue)
		for head < levelEnd {
			idx := queue[head]
			head++

			if idx == n - 1 {
				return steps
			}

			if next := idx + 1; next < n && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
			if prev := idx - 1; prev >= 0 && !visited[prev] {
				visited[prev] = true
				queue = append(queue, prev)
			}

			same := indicesByValue[arr[idx]]
			for _, j := range same {
				if !visited[j] {
					visited[j] = true
					queue = append(queue, j)
				}
			}
			delete(indicesByValue, arr[idx])
		}

		steps++
	}

	return -1
}

func main() {
	// result: 3
	// arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}

	// result: 0
	// arr := []int{7}

	// result: 1
	arr := []int{7, 6, 9, 6, 9, 6, 9, 7}

	// result:
	// arr := []int{}

	result := minJumps(arr)
	fmt.Printf("result = %v\n", result)
}
