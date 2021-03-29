package main
import (
	"fmt"
)

var direction [][]int = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func pacificAtlantic(matrix [][]int) [][]int {
	// later code assume matrix[0] exist
	if len(matrix) == 0 {
		return [][]int{}
	}

	pacificVisited := make([][]bool, len(matrix))
	atlanticVisited := make([][]bool, len(matrix))
	for i := range pacificVisited {
		pacificVisited[i] = make([]bool, len(matrix[0]))
		atlanticVisited[i] = make([]bool, len(matrix[0]))
	}

	// pacific reachable points (matrix left or top edge)
	q1 := make([][]int, 0)
	for j := range matrix[0] {
		q1 = append(q1, []int{0, j})
	}
	for i := 1; i < len(matrix); i++ {
		q1 = append(q1, []int{i, 0})
	}
	bfs(matrix, q1, pacificVisited)

	// atlantic reachable points (matrix right or bottom edge)
	q2 := make([][]int, 0)
	for j := range matrix[0] {
		q2 = append(q2, []int{len(matrix) - 1, j})
	}
	for i := 0; i < len(matrix)-1; i++ {
		q2 = append(q2, []int{i, len(matrix[0]) - 1})
	}
	bfs(matrix, q2, atlanticVisited)

	result := make([][]int, 0)

	for i := range pacificVisited {
		for j := range pacificVisited[0] {
			if pacificVisited[i][j] && atlanticVisited[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func bfs(matrix [][]int, queue [][]int, visited [][]bool) {
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if visited[q[0]][q[1]] {
			continue
		}

		for _, dir := range direction {
			x := q[1] + dir[1]
			y := q[0] + dir[0]

			if validPoint(matrix, x, y) && matrix[y][x] >= matrix[q[0]][q[1]] {
				queue = append(queue, []int{y, x})
			}
		}

		visited[q[0]][q[1]] = true
	}
}

func validPoint(matrix [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix[0]) && y < len(matrix)
}

// Wrong Answer
func ngSolution(matrix [][]int) [][]int {
    if len(matrix) <= 0 || len(matrix[0]) <= 0 {
        return [][]int{}
    }
    result := [][]int{}
    for row, nums := range matrix {
        for col, _ := range nums {
            if canFlowPacific(row, col, matrix) && canFlowAtlantic(row, col, matrix) {
                result = append(result, []int{row, col})
            }
        }
    }
    return result
}

func canFlowPacific(row int, col int, matrix [][]int) bool {
    currentXNum := matrix[row][col]
    currentYNum := matrix[row][col]
    canFlowX := true
    canFlowY := true
    savedXCol := int(-1)
    savedYRow := int(-1)
    for i := col - 1; i >= 0; i-- {
        num := matrix[row][i]
        if currentXNum >= num {
            currentXNum = num
            savedXCol = i
            continue
        }
        canFlowX = false
        break
    }
    if !canFlowX && savedXCol >= 0 {
        canFlowX = canFlowPacific(row, savedXCol, matrix)
    }
    for i := row - 1; i >= 0; i-- {
        num := matrix[i][col]
        if currentYNum >= num {
            currentYNum = num
            savedYRow = i
            continue
        }
        canFlowY = false
        break
    }
    if !canFlowX && !canFlowY && savedYRow >= 0 {
        canFlowY = canFlowPacific(savedYRow, col, matrix)
    }
    return canFlowX || canFlowY
}

func canFlowAtlantic(row int, col int, matrix [][]int) bool {
    maxX := len(matrix[0]) - 1
    maxY := len(matrix) - 1
    currentXNum := matrix[row][col]
    currentYNum := matrix[row][col]
    canFlowX := true
    canFlowY := true
    savedXCol := int(-1)
    savedYRow := int(-1)
    for i := col + 1; i <= maxX; i++ {
        num := matrix[row][i]
        if currentXNum >= num {
            currentXNum = num
            savedXCol = i
            continue
        }
        canFlowX = false
        break
    }
    if !canFlowX && savedXCol >= 0 {
        canFlowX = canFlowAtlantic(row, savedXCol, matrix)
    }
    for i := row + 1; i <= maxY; i++ {
        num := matrix[i][col]
        if currentYNum >= num {
            currentYNum = num
            savedYRow = i
            continue
        }
        canFlowY = false
        break
    }
    if !canFlowX && !canFlowY && savedYRow >= 0 {
        canFlowY = canFlowAtlantic(savedYRow, col, matrix)
    }
    // fmt.Printf("[%v, %v] can flow to atlantic x: %v | y: %v\n", row, col, canFlowX, canFlowY)
    return canFlowX || canFlowY
}

func main() {
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	// matrix := [][]int{
	// 	{1, 2},
	// 	{4, 3},
	// }

	// matrix := [][]int{
	// 	{1, 1},
	// 	{1, 1},
	// 	{1, 1},
	// }

	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{8, 9, 4},
	// 	{7, 6, 5},
	// }
	result := pacificAtlantic(matrix)
	fmt.Printf("result = %v\n", result)
}

