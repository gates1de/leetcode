package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X int
	Y int
}

func minCost(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	points := make([]Point, 0, m * n)
	for i := range m {
		for j := range n {
			points = append(points, Point{X: i, Y: j})
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return grid[points[i].X][points[i].Y] < grid[points[j].X][points[j].Y]
	})

	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt32
		}
	}

	for t := 0; t <= k; t++ {
		minCost := math.MaxInt32

		for i, j := 0, 0; i < len(points); i++ {
			minCost = min(minCost, costs[points[i].X][points[i].Y])

			if i + 1 < len(points) && grid[points[i].X][points[i].Y] == grid[points[i + 1].X][points[i + 1].Y] {
				continue
			}

			for r := j; r <= i; r++ {
				costs[points[r].X][points[r].Y] = minCost
			}

			j = i + 1
		}

		for i := m - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if i == m - 1 && j == n - 1 {
					costs[i][j] = 0
					continue
				}

				if i != m - 1 {
					costs[i][j] = min(costs[i][j], costs[i + 1][j] + grid[i + 1][j])
				}
				if j != n - 1 {
					costs[i][j] = min(costs[i][j], costs[i][j + 1] + grid[i][j + 1])
				}
			}
		}
	}

	return costs[0][0]
}

func main() {
	// result: 7
	// grid := [][]int{{1,3,3},{2,5,4},{4,3,5}}
	// k := int(2)

	// result: 9
	grid := [][]int{{1,2},{2,3},{3,4}}
	k := int(1)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := minCost(grid, k)
	fmt.Printf("result = %v\n", result)
}

