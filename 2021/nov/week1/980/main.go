package main
import (
	"fmt"
)

var res int

func uniquePathsIII(grid [][]int) int {
    cellsToVisit := 0
    res = 0

    si, sj := 0, 0
    for i, row := range grid {
        for j, v := range row {
            switch v {
                case 1:
                    si, sj = i, j
                case 0:
                    cellsToVisit++
            }
        }
    }

    moving(grid, sj, si, cellsToVisit + 1)
    return res
}

func moving(grid [][]int, x int, y int, cellsToVisit int) {
    if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) || grid[y][x] == -1 || grid[y][x] == 3 {
        return
    }
    if grid[y][x] == 2 {
        if cellsToVisit == 0 { res++ }
        return
    }

    grid[y][x] = 3
    moving(grid, x + 1, y, cellsToVisit - 1)
    moving(grid, x - 1, y, cellsToVisit - 1)
    moving(grid, x, y + 1, cellsToVisit - 1)
    moving(grid, x, y - 1, cellsToVisit - 1)
    grid[y][x] = 0
}

// Slow & Use more memory
func mySolution(grid [][]int) int {
	result := int(0)
	goalStep := int(0)
	start := [2]int{-1, -1}
	visited := map[[2]int]bool{}
	for i, _ := range grid {
		for j, cell := range grid[i] {
			if cell == 1 {
				start = [2]int{i, j}
			} else if cell == -1 {
				visited[[2]int{i, j}] = true
			} else {
				goalStep++
			}
		}
	}

	if start[0] == -1 || start[1] == -1 {
		return 0
	}

	// fmt.Printf("goalStep = %v, start = %v, visited = %v\n", goalStep, start, visited)
	helper(0, goalStep, start[0], start[1], visited, &result, grid)
	return result
}

func helper(step int, goalStep int, y int, x int, visited map[[2]int]bool, result *int, grid [][]int) {
	if visited[[2]int{y, x}] {
		return
	}

	// fmt.Printf("step = %v, y = %v, x = %v\n", step, y, x)
	if grid[y][x] == 2 {
		// fmt.Printf("step = %v, goalStep = %v\n", step, goalStep)
		if step == goalStep {
			*result++
		}
		return
	}

	visited[[2]int{y, x}] = true

	if y - 1 >= 0 && !visited[[2]int{y - 1, x}] {
		newVisited := copyMap(visited)
		helper(step + 1, goalStep, y - 1, x, newVisited, result, grid)
	}
	if x + 1 <= len(grid[0]) - 1 && !visited[[2]int{y, x + 1}] {
		newVisited := copyMap(visited)
		helper(step + 1, goalStep, y, x + 1, newVisited, result, grid)
	}
	if y + 1 <= len(grid) - 1 && !visited[[2]int{y + 1, x}] {
		newVisited := copyMap(visited)
		helper(step + 1, goalStep, y + 1, x, newVisited, result, grid)
	}
	if x - 1 >= 0 && !visited[[2]int{y, x - 1}] {
		newVisited := copyMap(visited)
		helper(step + 1, goalStep, y, x - 1, newVisited, result, grid)
	}
}

func copyMap(m map[[2]int]bool) map[[2]int]bool {
    result := map[[2]int]bool{}
    for k, v := range m {
        result[k] = v
    }
    return result
}

func main() {
	// result: 2
	// grid := [][]int{
	// 	{1,0,0,0},
	// 	{0,0,0,0},
	// 	{0,0,2,-1},
	// }

	// result: 4
	// grid := [][]int{
	// 	{1,0,0,0},
	// 	{0,0,0,0},
	// 	{0,0,0,2},
	// }

	// result: 0
	// grid := [][]int{{0, 1}, {2, 0}}

	// result: 1
	// grid := [][]int{{1}, {2}}

	// result: 0
	grid := [][]int{{2}}

	// result: 
	// grid := [][]int{}

	result := uniquePathsIII(grid)
	fmt.Printf("result = %v\n", result)
}

