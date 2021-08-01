package main
import (
	"fmt"
)

func largestIsland(grid [][]int) int {
	isAllOne := true
    result := int(0)
    for i, row := range grid {
        for j, cell := range row {
            if cell == 1 {
                continue
            }
			isAllOne = false
			grid[i][j] = 1
            areaOfIsland := traversal(i, j, grid)
            // fmt.Printf("[%v, %v]: area of island = %v\n", i, j, areaOfIsland)
            if result < areaOfIsland {
                result = areaOfIsland
            }
			grid[i][j] = 0
        }
    }
	if isAllOne {
		return len(grid) * len(grid[0])
	}
    return result
}

func traversal(i int, j int, grid [][]int) int {
    result := int(0)
    visited := make([][]int, len(grid))
    for i, row := range grid {
        visited[i] = make([]int, len(row))
    }
    queue := [][]int{{i, j}}

    for len(queue) > 0 {
        // fmt.Printf("queue = %v\n", queue)
        coord := queue[len(queue) - 1]
        if len(queue) == 1 {
            queue = [][]int{}
        } else {
            queue = queue[:len(queue) - 1]
        }

        y := coord[0]
        x := coord[1]
        if visited[y][x] == 0 && grid[y][x] == 1 {
            result += 1
        }
        if visited[y][x] == 1 {
            continue
        }
        visited[y][x] = 1

        maxY := len(grid)
        maxX := len(grid[0])
        // up
        if y - 1 >= 0 && grid[y - 1][x] == 1 {
            queue = append(queue, []int{y - 1, x})
        }
        // right
        if x + 1 < maxX && grid[y][x + 1] == 1 {
            queue = append(queue, []int{y, x + 1})
        }
        // down
        if y + 1 < maxY && grid[y + 1][x] == 1 {
            queue = append(queue, []int{y + 1, x})
        }
        // left
        if x - 1 >= 0 && grid[y][x - 1] == 1 {
            queue = append(queue, []int{y, x - 1})
        }
    }
    return result
}

func main() {
	// result: 3
	// grid := [][]int{{1, 0}, {0, 1}}

	// result: 4
	// grid := [][]int{{1, 1}, {1, 0}}

	// result: 4
	// grid := [][]int{{1, 1}, {1, 1}}

	// result: 10
	grid := [][]int{
		{1, 0, 0, 0, 1},
		{1, 1, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 0, 0, 1},
	}

	// result: 
	// grid := [][]int{}

	result := largestIsland(grid)
	fmt.Printf("result = %v\n", result)
}

