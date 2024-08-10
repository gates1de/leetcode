package main
import (
	"fmt"
)

func regionsBySlashes(grid []string) int {
	n := len(grid)
	pointsPerSide := n + 1
	totalPoints := pointsPerSide * pointsPerSide

	parents := make([]int, totalPoints)
	for i, _ := range parents {
		parents[i] = -1
	}

	for i := 0; i < pointsPerSide; i++ {
		for j := 0; j < pointsPerSide; j++ {
			if i == 0 || j == 0 || i == pointsPerSide - 1 || j == pointsPerSide - 1 {
				point := i * pointsPerSide + j
				parents[point] = 0
			}
		}
	}

	parents[0] = -1
	result := int(1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' {
				topRight := i * pointsPerSide + (j + 1)
				bottomLeft := (i + 1) * pointsPerSide + j
				result += union(parents, topRight, bottomLeft)
			} else if grid[i][j] == '\\' {
				topLeft := i * pointsPerSide + j
				bottomRight := (i + 1) * pointsPerSide + (j + 1)
				result += union(parents, topLeft, bottomRight)
			}
		}
	}

	return result
}

func find(parents []int, node int) int {
	if parents[node] == -1 {
		return node
	}
	parents[node] = find(parents, parents[node])
	return parents[node]
}

func union(parents []int, node1 int, node2 int) int {
	p1 := find(parents, node1)
	p2 := find(parents, node2)

	if p1 == p2 {
		return 1
	}

	parents[p2] = p1
	return 0
}

func main() {
	// result: 2
	// grid := []string{" /","/ "}

	// result: 1
	// grid := []string{" /","  "}

	// result: 5
	grid := []string{"/\\","\\/"}

	// result: 
	// grid := []string{}

	result := regionsBySlashes(grid)
	fmt.Printf("result = %v\n", result)
}

