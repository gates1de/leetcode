package main
import (
	"fmt"
)

func shortestBridge(grid [][]int) int {
	n := len(grid)
	firstX := int(-1)
	firstY := int(-1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				firstX = i
				firstY = j
				break
			}
		}
	}

	queue := make([][]int, 0)
	secondQueue := make([][]int, 0)
	queue = append(queue, []int{firstX, firstY})
	secondQueue = append(secondQueue, []int{firstX, firstY})
	grid[firstX][firstY] = 2

	for len(queue) > 0 {
		newQueue := make([][]int, 0)
		for _, cell := range queue {
			x := cell[0]
			y := cell[1]
			for _, next := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
				currentX := next[0]
				currentY := next[1]
				if 0 <= currentX && currentX < n && 0 <= currentY && currentY < n && grid[currentX][currentY] == 1 {
					newQueue = append(newQueue, []int{currentX, currentY})
					secondQueue = append(secondQueue, []int{currentX, currentY})
					grid[currentX][currentY] = 2
				}
			}
		}
		queue = newQueue
	}

	result := int(0)
	for len(secondQueue) > 0 {
		newQueue := make([][]int, 0)
		for _, cell := range secondQueue {
			x := cell[0]
			y := cell[1]
			for _, next := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
				currentX := next[0]
				currentY := next[1]
				if 0 <= currentX && currentX < n && 0 <= currentY && currentY < n {
					if grid[currentX][currentY] == 1 {
						return result
					} else if grid[currentX][currentY] == 0 {
						newQueue = append(newQueue, []int{currentX, currentY})
						grid[currentX][currentY] = -1
					}
				}
			}
		}
		secondQueue = newQueue
		result++
	}

	return result
}

func main() {
	// result: 1
	// grid := [][]int{{0,1},{1,0}}

	// result: 2
	// grid := [][]int{{0,1,0},{0,0,0},{0,0,1}}

	// result: 1
	grid := [][]int{{1,1,1,1,1},{1,0,0,0,1},{1,0,1,0,1},{1,0,0,0,1},{1,1,1,1,1}}

	// result: 
	// grid := [][]int{}

	result := shortestBridge(grid)
	fmt.Printf("result = %v\n", result)
}

