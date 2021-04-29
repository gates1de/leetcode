package main
import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	} else if len(obstacleGrid) == 1 {
		for _, num := range obstacleGrid[0] {
			if num == 1 {
				return 0
			}
		}
		return 1
	} else if len(obstacleGrid[0]) == 1 {
		for _, nums := range obstacleGrid {
			if nums[0] == 1 {
				return 0
			}
		}
		return 1
	}

	maxY := len(obstacleGrid) - 1
	maxX := len(obstacleGrid[0]) - 1
	if obstacleGrid[maxY][maxX] == 1 {
		return 0
	}
	result := int(0)
	move(0, 0, obstacleGrid, &result)
	return result
}

func move(currentX int, currentY int, grid [][]int, goalCount *int) {
	// fmt.Printf("current point = [%v, %v]\n", currentX, currentY)
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	if currentX == maxX && currentY == maxY {
		*goalCount++
		return
	}
	if grid[currentY][currentX] == 1 {
		return
	}
	if currentX < maxX {
		move(currentX + 1, currentY, grid, goalCount)
	}
	if currentY < maxY {
		move(currentX, currentY + 1, grid, goalCount)
	}
}

func main() {
	// result: 2
	// obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	// result: 1
	// obstacleGrid := [][]int{{0, 1}, {0, 0}}

	// result: 0
	obstacleGrid := [][]int{{0, 0}, {0, 1}}

	// result: 
	// obstacleGrid := [][]int{}

	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("result = %v\n", result)
}

