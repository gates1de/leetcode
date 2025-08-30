package main

import (
	"fmt"
)

func lenOfVDiagonal(grid [][]int) int {
	dirs := [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	m := len(grid)
	n := len(grid[0])
	memo := make([][][4][2]int, m)

	for i := range memo {
		memo[i] = make([][4][2]int, n)

		for j := range memo[i] {
			for k := range memo[i][j] {
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}

	result := int(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for direction := 0; direction < 4; direction++ {
					result = max(result, dfs(i, j, direction, true, 2, dirs, grid, memo) + 1)
				}
			}
		}
	}

	return result
}

func dfs(
	cx int,
	cy int,
	direction int,
	turn bool,
	target int,
	dirs [4][2]int,
	grid [][]int,
	memo [][][4][2]int,
) int {
	nx := cx + dirs[direction][0]
	ny := cy + dirs[direction][1]

	m := len(grid)
	n := len(grid[0])

	if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != target {
		return 0
	}

	turnInt := int(0)
	if turn {
		turnInt = 1
	}

	if memo[nx][ny][direction][turnInt] != -1 {
		return memo[nx][ny][direction][turnInt]
	}

	maxStep := dfs(nx, ny, direction, turn, 2 - target, dirs, grid, memo)
	if turn {
		maxStep = max(maxStep, dfs(nx, ny, (direction+1)%4, false, 2 - target, dirs, grid, memo))
	}

	memo[nx][ny][direction][turnInt] = maxStep + 1
	return maxStep + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// grid := [][]int{{2,2,1,2,2},{2,0,2,2,0},{2,0,1,1,0},{1,0,2,2,2},{2,0,0,2,2}}

	// result: 4
	// grid := [][]int{{2,2,2,2,2},{2,0,2,2,0},{2,0,1,1,0},{1,0,2,2,2},{2,0,0,2,2}}

	// result: 5
	// grid := [][]int{{1,2,2,2,2},{2,2,2,2,0},{2,0,0,0,0},{0,0,2,2,2},{2,0,0,2,0}}

	// result: 1
	grid := [][]int{{1}}

	// result: 
	// grid := [][]int{}

	result := lenOfVDiagonal(grid)
	fmt.Printf("result = %v\n", result)
}

