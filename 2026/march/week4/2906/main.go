package main

import (
	"fmt"
)

const modulo = 12345

func constructProductMatrix(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, m)
	}

	suffix := int64(1)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = int(suffix)
			suffix = (suffix * int64(grid[i][j])) % modulo
		}
	}

	prefix := int64(1)
	for i := range n {
		for j := range m {
			p[i][j] = int((int64(p[i][j]) * prefix) % modulo)
			prefix = (prefix * int64(grid[i][j])) % modulo
		}
	}

	return p
}

func main() {
	// result: [[24,12],[8,6]]
	// grid := [][]int{{1,2},{3,4}}

	// result: [[2],[0],[0]]
	grid := [][]int{{12345},{2},{1}}

	// result: []
	// grid := [][]int{}

	result := constructProductMatrix(grid)
	fmt.Printf("result = %v\n", result)
}

