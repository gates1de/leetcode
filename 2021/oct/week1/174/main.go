package main
import (
	"fmt"
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon) - 1
	n := len(dungeon[0]) - 1

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i == m && j == n {
				dungeon[i][j] = max(1, 1 - dungeon[i][j])
				continue
			}
			if i == m {
				dungeon[i][j] = max(1, dungeon[i][j + 1] - dungeon[i][j])
				continue
			}
			if j == n {
				dungeon[i][j] = max(1, dungeon[i + 1][j] - dungeon[i][j])
				continue
			}
			dungeon[i][j] = max(1, min(dungeon[i + 1][j], dungeon[i][j + 1]) - dungeon[i][j])
		}
	}

	return dungeon[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time Limit Exceeded
func ngSolution(dungeon [][]int) int {
	result := math.MinInt32
	helper(0, 0, dungeon[0][0], dungeon[0][0], dungeon, &result)
	if result >= 0 {
		return 1
	}
	return 1 - result
}

func helper(y int, x int, sumDamage int, min int, dungeon [][]int, result *int) {
	if y == len(dungeon) - 1 && x == len(dungeon[0]) - 1 {
		// fmt.Printf("min = %v\n", min)
		if *result < min {
			*result = min
		}
		return
	}

    if x + 1 < len(dungeon[0]) {
		damage := dungeon[y][x + 1]
        if damage >= 0 {
			helper(y, x + 1, sumDamage + damage, min, dungeon, result)
        } else {
			if min > sumDamage + damage {
				helper(y, x + 1, sumDamage + damage, sumDamage + damage, dungeon, result)
			} else {
				helper(y, x + 1, sumDamage + damage, min, dungeon, result)
			}
		}
    }
    // down
    if y + 1 < len(dungeon) {
		damage := dungeon[y + 1][x]
        if damage >= 0 {
			helper(y + 1, x, sumDamage + damage, min, dungeon, result)
        } else {
			if min > sumDamage + damage {
				helper(y + 1, x, sumDamage + damage, sumDamage + damage, dungeon, result)
			} else {
				helper(y + 1, x, sumDamage + damage, min, dungeon, result)
			}
		}
    }
}

func main() {
	// result: 7
	// dungeon := [][]int{
	// 	{-2, -3, 3},
	// 	{-5, -10, 1},
	// 	{10, 30, -5},
	// }

	// result: 1
	// dungeon := [][]int{{0}}

	// result: 6
	// dungeon := [][]int{
	// 	{-2, -3, 3},
	// 	{-3, -10, 1},
	// 	{10, 30, -5},
	// }

	// result: 201
	dungeon := [][]int{{-200}}

	// result: 
	// dungeon := [][]int{
	// }

	result := calculateMinimumHP(dungeon)
	fmt.Printf("result = %v\n", result)
}

