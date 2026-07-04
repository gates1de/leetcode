package main

import (
	"fmt"
)

const inf = int(1e9)

func findSafeWalk(grid [][]int, health int) bool {
	m := len(grid)
	n := len(grid[0])

	dist := make([][]int, m)
	for i := range m {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = inf
		}
	}

	dist[0][0] = grid[0][0]

	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	dirs := []int{0, 1, 0, -1, 0}
	for {
		bestX, bestY := -1, -1
		bestCost := inf
		for i := range m {
			for j := range n {
				if !visited[i][j] && dist[i][j] < bestCost {
					bestCost = dist[i][j]
					bestX = i
					bestY = j
				}
			}
		}

		if bestX == -1 {
			break
		}

		visited[bestX][bestY] = true
		for k := range 4 {
			nx := bestX + dirs[k]
			ny := bestY + dirs[k + 1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			nextCost := dist[bestX][bestY] + grid[nx][ny]
			if nextCost < dist[nx][ny] {
				dist[nx][ny] = nextCost
			}
		}
	}

	return dist[m - 1][n - 1] < health
}

func main() {
	// result: true
	// grid := [][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}}
	// health := int(1)

	// result: false
	// grid := [][]int{{0,1,1,0,0,0},{1,0,1,0,0,0},{0,1,1,1,0,1},{0,0,1,0,1,0}}
	// health := int(3)

	// result: true
	grid := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	health := int(5)

	// result: 
	// grid := [][]int{}
	// health := int(0)

	result := findSafeWalk(grid, health)
	fmt.Printf("result = %v\n", result)
}
