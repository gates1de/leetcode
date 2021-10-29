package main
import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	rotCount := int(0)
	queue := [][2]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				rotCount++
			} else if cell == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// fmt.Printf("rotCount = %v, queue = %v\n", rotCount, queue)
	if rotCount == 0 {
		return 0
	}
	if len(queue) == 0 {
		if rotCount > 0 {
			return -1
		}
		return 0
	}

	result := int(0)
	for rotCount > 0 {
		newQueue := [][2]int{}
		result++

		for i := 0; i < len(queue); i++ {
			cell := queue[i]
			y, x := cell[0], cell[1]

			// up
			if y - 1 >= 0 && grid[y - 1][x] == 1 {
				grid[y - 1][x] = 2
				newQueue = append(newQueue, [2]int{y - 1, x})
				rotCount--
			}

			// right
			if x + 1 < len(grid[0]) && grid[y][x + 1] == 1 {
				grid[y][x + 1] = 2
				newQueue = append(newQueue, [2]int{y, x + 1})
				rotCount--
			}

			// down
			if y + 1 < len(grid) && grid[y + 1][x] == 1 {
				grid[y + 1][x] = 2
				newQueue = append(newQueue, [2]int{y + 1, x})
				rotCount--
			}

			// left
			if x - 1 >= 0 && grid[y][x - 1] == 1 {
				grid[y][x - 1] = 2
				newQueue = append(newQueue, [2]int{y, x - 1})
				rotCount--
			}
		}

		if len(newQueue) == 0 {
			break
		}
		queue = newQueue
	}

	if rotCount > 0 {
		return -1
	}

	return result
}

func main() {
	// result: 4
	// grid := [][]int{
	// 	{2, 1, 1},
	// 	{1, 1, 0},
	// 	{0, 1, 1},
	// }

	// result: -1
	// grid := [][]int{
	// 	{2, 1, 1},
	// 	{0, 1, 1},
	// 	{1, 0, 1},
	// }

	// result: 0
	// grid := [][]int{{0, 2}}

	// result: -1
	grid := [][]int{{1}}

	// result: 
	// grid := [][]int{
	// 	{0, 0, 0},
	// }

	result := orangesRotting(grid)
	fmt.Printf("result = %v\n", result)
}

