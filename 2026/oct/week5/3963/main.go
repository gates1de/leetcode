package main

import (
	"fmt"
)

func createGrid(m int, n int) []string {
	if m <= 0 || n <= 0 {
		return []string{}
	}

	grid := make([]string, m)
	for i := range grid {
		row := make([]byte, n)
		for j := range row {
			row[j] = '#'
		}
		if i == 0 {
			for j := range row {
				row[j] = '.'
			}
		} else {
			row[n-1] = '.'
		}
		grid[i] = string(row)
	}

	return grid
}

func main() {
	// result: ["..#","#.."]
	// m := int(2)
	// n := int(3)

	// result: ["..#","#..","##."]
	// m := int(3)
	// n := int(3)

	// result: ["...."]
	m := int(1)
	n := int(4)

	// result: []
	// m := int(0)
	// n := int(0)

	result := createGrid(m, n)
	fmt.Printf("result = %v\n", result)
}
