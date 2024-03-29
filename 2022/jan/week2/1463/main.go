package main
import (
	"fmt"
)

func cherryPickup(grid [][]int) int {
    dp := make([][][]int, len(grid))
    for i := range dp {
        columns := make([][]int, len(grid[0]))
        for j := range columns {
            column := make([]int, len(grid[0]))
            for x := range column {
                column[x] = -1
            }
            columns[j] = column

        }
        dp[i]= columns
    }
    return dfs(grid, 0, 0, len(grid[0]) - 1, dp)
}

func dfs(grid [][]int, row int, col1 int, col2 int, dp [][][]int) int {
    if row == len(grid) {
        return 0
    }

    if dp[row][col1][col2] != -1 {
        return dp[row][col1][col2]
    }

    for i := -1; i <= 1; i++ {
        newCol1 := col1 + i
        for j := -1; j <=1; j++ {
            newCol2 := col2 + j
            if newCol1 >= 0 && newCol2 >= 0 && newCol1 < len(grid[0]) && newCol2 < len(grid[0]) {
                dp[row][col1][col2] = max(dp[row][col1][col2], dfs(grid, row + 1, newCol1, newCol2, dp))
            }
        }
    }

    if col1 == col2 {
        dp[row][col1][col2] += grid[row][col1]
    } else {
        dp[row][col1][col2] += grid[row][col1] + grid[row][col2]
    }
    return dp[row][col1][col2]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Time Limit Exceeded
func ngSolution(grid [][]int) int {
	dp := map[[3]int]int{}
	r1Start := []int{0, 0}
	r2Start := []int{0, len(grid[0]) - 1}
	r1StartCherry := grid[0][0]
	r2StartCherry := grid[0][len(grid[0]) - 1]
	dp[[3]int{0, 0, r2Start[1]}] = r1StartCherry + r2StartCherry
	helper(r1StartCherry + r2StartCherry, r1Start, r2Start, dp, grid)

	result := int(0)
	for cell, val := range dp {
		if cell[0] != len(grid) - 1 {
			continue
		}

		if result < val {
			result = val
		}
	}
	return result
}

func helper(currentCherry int, r1Cell []int, r2Cell []int, dp map[[3]int]int, grid [][]int) {
	// fmt.Printf("currentCherry = %v, r1Cell = %v, r2Cell = %v\n", currentCherry, r1Cell, r2Cell)
	y := r1Cell[0]
	if y >= len(grid) || r1Cell[1] == r2Cell[1] {
		return
	}

	if currentCherry < dp[[3]int{y, r1Cell[1], r2Cell[1]}] {
		return
	}

	dp[[3]int{y, r1Cell[1], r2Cell[1]}] = currentCherry

	maxX := len(grid[0]) - 1

	nextY := y + 1
	if nextY >= len(grid) {
		return
	}

	r1Left := r1Cell[1] - 1
	r1Below := r1Cell[1]
	r1Right := r1Cell[1] + 1
	r2Left := r2Cell[1] - 1
	r2Below := r2Cell[1]
	r2Right := r2Cell[1] + 1

	if r1Left >= 0 {
		// r1 -> left, r2 -> left
		if r2Left >= 0 {
			r1Cherry := grid[nextY][r1Left]
			r2Cherry := grid[nextY][r2Left]
			helper(
				currentCherry + r1Cherry + r2Cherry,
				[]int{nextY, r1Left},
				[]int{nextY, r2Left},
				dp,
				grid,
			)
		}

		// r1 -> left, r2 -> below
		r1Cherry := grid[nextY][r1Left]
		r2Cherry := grid[nextY][r2Below]
		helper(
			currentCherry + r1Cherry + r2Cherry,
			[]int{nextY, r1Left},
			[]int{nextY, r2Below},
			dp,
			grid,
		)

		// r1 -> left, r2 -> right
		if r2Right <= maxX {
			r1Cherry := grid[nextY][r1Left]
			r2Cherry := grid[nextY][r2Right]
			helper(
				currentCherry + r1Cherry + r2Cherry,
				[]int{nextY, r1Left},
				[]int{nextY, r2Right},
				dp,
				grid,
			)
		}
	}

	// r1 -> below, r2 -> left
	if r2Left >= 0 {
		r1Cherry := grid[nextY][r1Below]
		r2Cherry := grid[nextY][r2Left]
		helper(
			currentCherry + r1Cherry + r2Cherry,
			[]int{nextY, r1Below},
			[]int{nextY, r2Left},
			dp,
			grid,
		)
	}

	// r1 -> below, r2 -> below
	r1Cherry := grid[nextY][r1Below]
	r2Cherry := grid[nextY][r2Below]
	helper(
		currentCherry + r1Cherry + r2Cherry,
		[]int{nextY, r1Below},
		[]int{nextY, r2Below},
		dp,
		grid,
	)

	// r1 -> below, r2 -> right
	if r2Right <= maxX {
		r1Cherry := grid[nextY][r1Below]
		r2Cherry := grid[nextY][r2Right]
		helper(
			currentCherry + r1Cherry + r2Cherry,
			[]int{nextY, r1Below},
			[]int{nextY, r2Right},
			dp,
			grid,
		)
	}

	if r1Right <= maxX {
		// r1 -> right, r2 -> left
		if r2Left >= 0 {
			r1Cherry := grid[nextY][r1Right]
			r2Cherry := grid[nextY][r2Left]
			helper(
				currentCherry + r1Cherry + r2Cherry,
				[]int{nextY, r1Right},
				[]int{nextY, r2Left},
				dp,
				grid,
			)
		}

		// r1 -> right, r2 -> below
		r1Cherry := grid[nextY][r1Right]
		r2Cherry := grid[nextY][r2Below]
		helper(
			currentCherry + r1Cherry + r2Cherry,
			[]int{nextY, r1Right},
			[]int{nextY, r2Below},
			dp,
			grid,
		)

		// r1 -> right, r2 -> right
		if r2Right <= maxX {
			r1Cherry := grid[nextY][r1Right]
			r2Cherry := grid[nextY][r2Right]
			helper(
				currentCherry + r1Cherry + r2Cherry,
				[]int{nextY, r1Right},
				[]int{nextY, r2Right},
				dp,
				grid,
			)
		}
	}
}

func main() {
	// result: 24
	// grid := [][]int{
	// 	{3,1,1},
	// 	{2,5,1},
	// 	{1,5,5},
	// 	{2,1,1},
	// }

	// result: 28
	grid := [][]int{
		{1,0,0,0,0,0,1},
		{2,0,0,0,0,3,0},
		{2,0,9,0,0,0,0},
		{0,3,0,5,4,0,0},
		{1,0,2,3,0,0,6},
	}

	// result: 
	// grid := [][]int{}

	result := cherryPickup(grid)
	fmt.Printf("result = %v\n", result)
}

