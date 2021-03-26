package main
import (
	"fmt"
)

func pacificAtlantic(matrix [][]int) [][]int {
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

