package main
import (
	"fmt"
)

func uniquePaths(m int, n int) int {
    visited := make([][]int, m)
    for i, _ := range visited {
        visited[i] = make([]int, n)
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            visited[i][j] = helper(i, j, 0, m, n, visited)
        }
    }

    visited[m - 1][n - 1] = 1

    return visited[0][0]
}

func helper(i int, j int, count int, m int, n int, visited [][]int) int {
    if i == m - 1 && j == n - 1 {
        return count
    }

    if visited[i][j] > 0 {
        return count * visited[i][j]
    }

    newCount := int(0)
    if i + 1 < m {
        newCount += helper(i + 1, j, count + 1, m, n, visited)
    }
    if j + 1 < n {
        newCount += helper(i, j + 1, count + 1, m, n, visited)
    }

    return newCount
}

func main() {
	// result: 28
	// m := int(3)
	// n := int(7)

	// result: 3
	m := int(3)
	n := int(2)

	// result: 
	// m := int(0)
	// n := int(0)

	result := uniquePaths(m, n)
	fmt.Printf("result = %v\n", result)
}

