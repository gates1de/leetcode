package main
import (
	"fmt"
	"math"
)

func nearestExit(maze [][]byte, entrance []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}

	m := len(maze)
	n := len(maze[0])
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	startX := entrance[1]
	startY := entrance[0]
	maze[startY][startX] = '+'

	queue := make([][]int, 0, m * n)
	queue = append(queue, []int{startY, startX, 0})

	for len(queue) > 0 {
		current := queue[0]
		currentX := current[1]
		currentY := current[0]
		currentDistance := current[2]
		queue = queue[1:]

		for _, dir := range dirs {
			nextX := currentX + dir[0]
			nextY := currentY + dir[1]

			if isStop(nextX, nextY, maze) {
				continue
			}

			if isGoal(nextX, nextY, maze) {
				return currentDistance + 1
			}

			maze[nextY][nextX] = '+'
			queue = append(queue, []int{nextY, nextX, currentDistance + 1})
		}
	}

	return -1
}
        
// Wrong Answer
func ngSolution(maze [][]byte, entrance []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}

	m := len(maze)
	n := len(maze[0])
	result := math.MaxInt32

	// -1 == init position, 0 == not visited, 1 == visited
	visited := make([][]int, m)
	for i, _ := range visited {
		visited[i] = make([]int, n)
	}

	initX := entrance[1]
	initY := entrance[0]

	if isStop(initX, initY, maze) {
		return -1
	}
	visited[initY][initX] = -1

	helper(0, initX, initY, maze, visited, &result)

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func helper(count int, x int, y int, maze [][]byte, visited [][]int, result *int) {
	if *result <= count {
		return
	}

	if visited[y][x] == 1 {
		return
	}

	if visited[y][x] == 0 {
		if isGoal(x, y, maze) {
			*result = count
			return
		}
		
		visited[y][x] = 1
	}

	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, dir := range dirs {
		newX := x + dir[0]
		newY := y + dir[1]
		if isStop(newX, newY, maze) {
			continue
		}
		helper(count + 1, newX, newY, maze, visited, result)
	}
}

func isGoal(x int, y int, maze [][]byte) bool {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return false
	}

	m := len(maze)
	n := len(maze[0])

	return x == 0 || x == n - 1 || y == 0 || y == m - 1
}

func isStop(x int, y int, maze [][]byte) bool {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return true
	}

	m := len(maze)
	n := len(maze[0])

	return x < 0 || x >= n || y < 0 || y >= m || maze[y][x] == '+'
}

func main() {
	// result: 1
	// maze := [][]byte{{'+','+','.','+'},{'.','.','.','+'},{'+','+','+','.'}}
	// entrance := []int{1,2}

	// result: 2
	// maze := [][]byte{{'+','+','+'},{'.','.','.'},{'+','+','+'}}
	// entrance := []int{1,0}

	// result: -1
	// maze := [][]byte{{'.','+'}}
	// entrance := []int{0,0}

	// result: -1
	maze := [][]byte{{'+','.'}}
	entrance := []int{0,1}

	// result: 
	// maze := [][]byte{}
	// entrance := []int{}

	result := nearestExit(maze, entrance)
	fmt.Printf("result = %v\n", result)
}

