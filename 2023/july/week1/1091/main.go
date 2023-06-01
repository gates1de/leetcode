package main
import (
	"fmt"
)

func getDirs() [][2]int {
    return [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
}

func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    if grid[0][0] == 1 || grid[n - 1][n - 1] == 1 {
        return -1
    }
    if n == 1 {
        return 1
    }

    visited := make([][]int, n)
    for i, _ := range visited {
        visited[i] = make([]int, n)
    }
    bfs([][]int{{0, 0, 1}}, grid, visited)

    if visited[n - 1][n - 1] == 0 {
        return -1
    }
    return visited[n - 1][n - 1]
}

func bfs(queue [][]int, grid [][]int, visited [][]int) {
    n := len(grid)
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]

        if visited[q[0]][q[1]] > 0 {
            continue
        }

        for _, dir := range getDirs() {
            x := q[1] + dir[1]
            y := q[0] + dir[0]

            if validPoint(x, y, n) && grid[y][x] == 0 {
                queue = append(queue, []int{y, x, q[2] + 1})
            }
        }

        visited[q[0]][q[1]] = q[2]
    }
}

func validPoint(x int, y int, n int) bool {
    return x >= 0 && y >= 0 && x < n && y < n
}

func main() {
	// result: 2
	// grid := [][]int{{0,1},{1,0}}

	// result: 4
	// grid := [][]int{{0,0,0},{1,1,0},{1,1,0}}

	// result: -1
	grid := [][]int{{1,0,0},{1,1,0},{1,1,0}}

	// result: 
	// grid := [][]int{}

	result := shortestPathBinaryMatrix(grid)
	fmt.Printf("result = %v\n", result)
}

