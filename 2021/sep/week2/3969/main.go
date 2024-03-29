package main
import (
	"fmt"
)

func orderOfLargestPlusSign(n int, mines [][]int) int {
    grid := make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            grid[i][j] = n
        }
    }
    for _, v := range mines {
        grid[v[0]][v[1]] = 0
    }
    for i := 0; i < n; i++ {
        left := 0
        right := 0
        up := 0
        down := 0
        for j, k := 0, n - 1; j < n; j, k = j + 1, k - 1 {
            if grid[i][j] == 0 {
                left = 0
            } else {
                left += 1
            }
            if left < grid[i][j] { grid[i][j] = left }
            if grid[i][k] == 0 {
                right = 0
            } else {
                right += 1
            }
            if right < grid[i][k] { grid[i][k] = right }
            if grid[j][i] == 0 {
                up = 0
            } else {
                up += 1
            }
            if up < grid[j][i] { grid[j][i] = up }
            if grid[k][i] == 0 {
                down = 0
            } else {
                down += 1
            }
            if down < grid[k][i] { grid[k][i] = down }
        }
    }
    res := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] > res { res = grid[i][j] }
        }
    }
    return res
}

// Time Limit Exceeded
func ngSolution(n int, mines [][]int) int {
	if n * n <= len(mines) {
		return 0
	}

	m := map[[2]int]bool{}
	for _, mine := range mines {
		m[[2]int{mine[0], mine[1]}] = true
	}

	result := int(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[[2]int{i, j}] {
				// fmt.Printf("[%v, %v] is 0\n", i, j)
				continue
			}

			vCount := int(1)
			up := i - 1
			down := i + 1
			// fmt.Printf("----------------- vCount -----------------\n")
			for k := 0; k < min(i, n - i - 1); k++ {
				if m[[2]int{up, j}] || m[[2]int{down, j}] {
					break
				}
				// fmt.Printf("[%v][%v] & [%v][%v]\n", up, j, down, j)
				up--
				down++
				vCount++
			}

			hCount := int(1)
			left := j - 1
			right := j + 1
			// fmt.Printf("----------------- hCount -----------------\n")
			for k := 0; k < min(j, n - j - 1); k++ {
				if m[[2]int{i, left}] || m[[2]int{i, right}] {
					break
				}
				// fmt.Printf("[%v][%v] & [%v][%v]\n", i, left, i, right)
				left--
				right++
				hCount++
			}

			size := min(vCount, hCount)
			if result < size {
				result = size
			}
		}
	}

    return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// n := int(5)
	// mines := [][]int{{4, 2}}

	// result: 0
	// n := int(1)
	// mines := [][]int{{0, 0}}

	// result: 3
	// n := int(5)
	// mines := [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}}

	// result: 1
	n := int(5)
	mines := [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}

	// result: 0
	// n := int(2)
	// mines := [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 0}}

	// result: 
	// n := int(0)
	// mines := [][]int{}

	result := orderOfLargestPlusSign(n, mines)
	fmt.Printf("result = %v\n", result)
}

