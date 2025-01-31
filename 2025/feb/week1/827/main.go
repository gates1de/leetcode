package main
import (
	"fmt"
	"math"
)

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	potIsland := make(map[int]int)
	islands := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			key := potKey(i, j)
			if _, ok := potIsland[key]; !ok {
				islands = dfs(grid, i, j, potIsland, islands)
			}
		}
	}
	max := maxInArr(islands)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			sum := countSum(grid, i, j, potIsland, islands)
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func potKey(i int, j int) int {
	return i * 500 + j
}

func dfsInt(islandInd int, grid [][]int, i int, j int, potIsland map[int]int) int{
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	key := potKey(i, j)
	if _, ok := potIsland[key]; ok {
		return 0
	}
	island := 1
	potIsland[key] = islandInd
	island += dfsInt(islandInd, grid, i - 1, j, potIsland)
	island += dfsInt(islandInd, grid, i, j - 1, potIsland)
	island += dfsInt(islandInd, grid, i, j + 1, potIsland)
	island += dfsInt(islandInd, grid, i + 1, j, potIsland)
	return island
}

func dfs(grid [][]int, i int, j int, potIsland map[int]int, islands []int) []int {
	islandInd := len(islands)
	oneIsland := dfsInt(islandInd, grid, i, j, potIsland)
	if oneIsland == 0 {
		return islands
	}
	return append(islands, oneIsland)

}

func maxInArr(arr []int) int {
	max := math.MinInt32
	for _, i := range(arr) {
		if max < i {
			max = i
		}
	}
	return max
}

func adjIsland(grid [][]int, i int, j int, potIsland map[int]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return -1
	}
	if island, ok := potIsland[potKey(i, j)]; ok {
		return island
	}
	return -1
}

func countSum(grid [][]int, i int, j int, potIsland map[int]int, islands []int) int {
	adjIslands := make(map[int]bool)
	if n := adjIsland(grid, i - 1, j, potIsland); n != -1 {
		adjIslands[n] = true
	}
	if n := adjIsland(grid, i, j - 1, potIsland); n != -1 {
		adjIslands[n] = true
	}
	if n := adjIsland(grid, i, j + 1, potIsland); n != -1 {
		adjIslands[n] = true
	}
	if n := adjIsland(grid, i + 1, j, potIsland); n != -1 {
		adjIslands[n] = true
	}
	sum := 1
	for k, _ := range(adjIslands) {
		sum += islands[k]
	}
	return sum
}

// Time Limit Exceeded
func ngSolution(grid [][]int) int {
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
