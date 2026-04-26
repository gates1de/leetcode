package main

import (
	"fmt"
)

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(grid [][]byte, x, y int) bool {
	n := len(grid)
	m := len(grid[0])
	return x >= 0 && x < n && y >= 0 && y < m
}

func containsCycle(grid [][]byte) bool {
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	for i := range n {
		for j := range m {
			if !visited[i][j] && dfs(grid, &visited, i, j, -1, -1) {
				return true
			}
		}
	}

	return false
}

func dfs(grid [][]byte, visited *[][]bool, i, j, x, y int) bool {
	if !isValid(grid, i, j) {
		return false
	}

	if (*visited)[i][j] {
		return true
	}

	(*visited)[i][j] = true

	found := false
	for _, d := range dirs {
		newX := i + d[0]
		newY := j + d[1]

		if isValid(grid, newX, newY) {
			if grid[newX][newY] == grid[i][j] && !(x == newX && y == newY) {
				found = found || dfs(grid, visited, newX, newY, i, j)

				if found {
					return found
				}
			}
		}
	}

	return found
}

func main() {
	// result: true
	// grid := [][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}}

	// result: true
	// grid := [][]byte{{'c','c','c','a'},{'c','d','c','c'},{'c','c','e','c'},{'f','c','c','c'}}

	// result: false
	grid := [][]byte{{'a','b','b'},{'b','z','b'},{'b','b','a'}}

	// result: 
	// grid := [][]byte{}

	result := containsCycle(grid)
	fmt.Printf("result = %v\n", result)
}

