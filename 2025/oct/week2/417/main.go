package main

import (
	"fmt"
)


var direction [][]int = [][]int{
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
}

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
    if m == 0 {
        return [][]int{}
    }
	n := len(heights[0])

    pacificVisited := make([][]bool, m)
    atlanticVisited := make([][]bool, m)
    for i := range pacificVisited {
        pacificVisited[i] = make([]bool, n)
        atlanticVisited[i] = make([]bool, n)
    }

    q1 := make([][]int, 0)
    for j := range heights[0] {
        q1 = append(q1, []int{0, j})
    }
    for i := 1; i < m; i++ {
        q1 = append(q1, []int{i, 0})
    }
    bfs(heights, q1, pacificVisited)

    q2 := make([][]int, 0)
    for j := range heights[0] {
        q2 = append(q2, []int{m - 1, j})
    }
    for i := 0; i < m - 1; i++ {
        q2 = append(q2, []int{i, n - 1})
    }
    bfs(heights, q2, atlanticVisited)

    result := make([][]int, 0)

    for i := range pacificVisited {
        for j := range pacificVisited[0] {
            if pacificVisited[i][j] && atlanticVisited[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func bfs(heights [][]int, queue [][]int, visited [][]bool) {
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]

        if visited[q[0]][q[1]] {
            continue
        }

        for _, dir := range direction {
            x := q[1] + dir[1]
            y := q[0] + dir[0]

            if validPoint(heights, x, y) && heights[y][x] >= heights[q[0]][q[1]] {
                queue = append(queue, []int{y, x})
            }
        }

        visited[q[0]][q[1]] = true
    }
}

func validPoint(heights [][]int, x, y int) bool {
    return x >= 0 && y >= 0 && x < len(heights[0]) && y < len(heights)
}

func main() {
	// result: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
	// heights := [][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}

	// result: [[0,0]]
	// heights := [][]int{{1}}

	// result: [[0,1],[1,0],[1,1]]
	// heights := [][]int{{1,2},{4,3}}

	// result: [[0,0],[0,1],[1,0],[1,1],[2,0],[2,1]]
	heights := [][]int{{1,1},{1,1},{1,1}}

	// result: 
	// heights := [][]int{}

	result := pacificAtlantic(heights)
	fmt.Printf("result = %v\n", result)
}

