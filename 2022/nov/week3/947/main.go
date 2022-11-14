package main
import (
	"fmt"
)

func removeStones(stones [][]int) int {
    rows := make(map[int][]int)
    cols := make(map[int][]int)
    
    for _, stone := range stones {
        rows[stone[0]] = append(rows[stone[0]], stone[1])
        cols[stone[1]] = append(cols[stone[1]], stone[0])
    }
    
    removeCount := int(0)
    visited := make(map[[2]int]bool)
    
    for _, stone := range stones {
        if _, ok := visited[[2]int{stone[0], stone[1]}]; !ok {
            removeCount++
            dfs(stone[0], stone[1], rows, cols, visited)
        } 
    }

    return len(stones) - removeCount
}

func dfs(x int, y int, rows map[int][]int, cols map[int][]int, visited map[[2]int]bool) {
    visited[[2]int{x, y}] = true

    for _, col := range rows[x] {
        if !visited[[2]int{x, col}] {
            dfs(x, col, rows, cols, visited)
        }        
    }

    for _, row := range cols[y] {
        if !visited[[2]int{row, y}] {
            dfs(row, y, rows, cols, visited)
        }        
    }
}

func main() {
	// result: 5
	// stones := [][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}}

	// result: 3
	// stones := [][]int{{0,0},{0,2},{1,1},{2,0},{2,2}}

	// result: 0
	stones := [][]int{{0, 0}}

	// result: 
	// stones := [][]int{}

	result := removeStones(stones)
	fmt.Printf("result = %v\n", result)
}

