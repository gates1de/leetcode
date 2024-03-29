package main
import (
	"fmt"
	"strconv"
)

type node struct {
    pos []int
}

func shortestPath(grid [][]int, k int) int {
    queue := []*node{{pos: []int{0, 0}}, nil}

    m,n  := len(grid), len(grid[0])
    leftK := make([][]int, m)
    for i := range leftK {
        leftK[i] = make([]int, n)
        for j := range leftK[i] {
            leftK[i][j] = -1
        }
    }

    leftK[0][0] = k
    var step int
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]
        if q == nil {
            step++
            if len(queue) != 0 {
                queue = append(queue, nil)
            }
            continue
        }

        x, y := q.pos[0], q.pos[1]
        if x == m-1 && y == n-1 {
            return step
        }

        currK := leftK[x][y]
        for _, d := range [][]int{{0,1}, {0, -1}, {1,0}, {-1, 0}} {
            qn := update(grid, leftK, x+d[0], y+d[1], currK)
            if qn != nil {
                queue = append(queue, &node{pos: qn})
            }
        }
    }

    return -1
}

func update(grid, leftK [][]int, x, y, k int) []int {
    if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0 {
        return nil
    }

    if grid[x][y] == 1 {
        k--
    }

    if k < 0 {
        return nil
    }

    if leftK[x][y] == -1 { // means not visited yet
        leftK[x][y] = k
        return []int{x, y}
    } else if leftK[x][y] < k { // means visited
        leftK[x][y] = k
        return []int{x, y}
    }

    return nil
}

// REF: https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/discuss/452124/Go-Simple-BFS-Solution
func shortestPath2(grid [][]int, k int) int {
    move := []int{0, 1, 0, -1, 0}
    queue:= make([][]int, 0, 0)
    queue = append(queue, []int{0, 0, k})
    set := make(map[string]bool)
    set["0" + " " + "0" + " " + strconv.Itoa(k)] = true
    step := 0
    m := len(grid)
    n := len(grid[0])
    for len(queue) > 0 {
        size:= len(queue)
        for i := 0; i < size; i++ {
            cur:= queue[i]
            if cur[0] == m - 1 && cur[1] == n - 1 {
                return step
            }
            for i:= 0; i < 4; i++{
                nx := cur[0] + move[i]
                ny := cur[1] + move[i+1]
                nk := cur[2]
                if nx >= 0 && nx < m && ny >= 0 && ny < n {
                    if grid[nx][ny] == 1 {
                        nk--
                    }
                    if nk < 0 {
                        continue
                    }
                    status:= strconv.Itoa(nx) + " " + strconv.Itoa(ny) + " " + strconv.Itoa(nk)
                    if _, ok := set[status]; !ok{
                        set[status] = true
                        queue = append(queue, []int{nx, ny, nk})
                    }
                }
            }
        }
        queue = queue[size:]
        step++
    }
    return -1
}

// Time Limit Exceeded
func ngSolution(grid [][]int, k int) int {
	result := int(10000)
	visited := map[[2]int]int{}
	for i, _ := range grid {
		for j, _ := range grid[i] {
			visited[[2]int{i, j}] = 10000
		}
	}
	helper(0, 0, 0, k, visited, grid, &result)
	if result == 10000 {
		return -1
	}
	return result
}

func helper(y int, x int, step int, k int, visited map[[2]int]int, grid [][]int, result *int) {
	// fmt.Printf("[%v, %v]\n", y, x)
	if visited[[2]int{y, x}] < step {
		return
	}
	if y == len(grid) - 1 && x == len(grid[0]) - 1 {
		// fmt.Printf("step = %v\n", step)
		if *result > step {
			*result = step
		}
		return
	}

	visited[[2]int{y, x}] = step
	// up
	if y - 1 >= 0 {
		if grid[y - 1][x] == 0 {
			helper(y - 1, x, step + 1, k, visited, grid, result)
		} else if grid[y - 1][x] == 1 && k > 0 {
			helper(y - 1, x, step + 1, k - 1, visited, grid, result)
		}
	}
	// right
	if x + 1 < len(grid[0]) {
		if grid[y][x + 1] == 0 {
			helper(y, x + 1, step + 1, k, visited, grid, result)
		} else if grid[y][x + 1] == 1 && k > 0 {
			helper(y, x + 1, step + 1, k - 1, visited, grid, result)
		}
	}
	// down
	if y + 1 < len(grid) {
		if grid[y + 1][x] == 0 {
			helper(y + 1, x, step + 1, k, visited, grid, result)
		} else if grid[y + 1][x] == 1 && k > 0 {
			helper(y + 1, x, step + 1, k - 1, visited, grid, result)
		}
	}
	// left
	if x - 1 >= 0 {
		if grid[y][x - 1] == 0 {
			helper(y, x - 1, step + 1, k, visited, grid, result)
		} else if grid[y][x - 1] == 1 && k > 0 {
			helper(y, x - 1, step + 1, k - 1, visited, grid, result)
		}
	}
}

func main() {
	// result: 6
	// grid := [][]int{
	// 	{0, 0, 0},
	// 	{1, 1, 0},
	// 	{0, 0, 0},
	// 	{0, 1, 1},
	// 	{0, 0, 0},
	// }
	// k := int(1)

	// result: -1
	// grid := [][]int{
	// 	{0, 1, 1},
	// 	{1, 1, 1},
	// 	{1, 0, 0},
	// }
	// k := int(1)

	// result: 0
	// grid := [][]int{{0}}
	// k := int(1)

	// result: 20
	// grid := [][]int{
	// 	{0,0,0,0,0,0,0,0,0,0},
	// 	{0,1,1,1,1,1,1,1,1,0},
	// 	{0,1,0,0,0,0,0,0,0,0},
	// 	{0,1,0,1,1,1,1,1,1,1},
	// 	{0,1,0,0,0,0,0,0,0,0},
	// 	{0,1,1,1,1,1,1,1,1,0},
	// 	{0,1,0,0,0,0,0,0,0,0},
	// 	{0,1,0,1,1,1,1,1,1,1},
	// 	{0,1,0,1,1,1,1,0,0,0},
	// 	{0,1,0,0,0,0,0,0,1,0},
	// 	{0,1,1,1,1,1,1,0,1,0},
	// 	{0,0,0,0,0,0,0,0,1,0},
	// }
	// k := int(1)

	// result: 387
	grid := [][]int{
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	}
	k := int(5)

	// result: 8
	// grid := [][]int{
	// 	{0,1,0,0,0,0},
	// 	{0,1,0,1,1,0},
	// 	{0,1,0,1,1,0},
	// 	{0,1,0,1,1,0},
	// }
	// k := int(1)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := shortestPath(grid, k)
	fmt.Printf("result = %v\n", result)
}

