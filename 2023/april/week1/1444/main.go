package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func ways(pizza []string, k int) int {
	rows := len(pizza)
	cols := len(pizza[0])
	apples := make([][]int, rows + 1)
	for i, _ := range apples {
		apples[i] = make([]int, cols + 1)
	}

	f := make([][]int, rows)
	for i, _ := range f {
		f[i] = make([]int, cols)
	}

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			tmp := int(0)
			if pizza[row][col] == 'A' {
				tmp = 1
			}
			apples[row][col] = tmp + apples[row + 1][col] + apples[row][col + 1] - apples[row + 1][col + 1]
			f[row][col] = 0
			if apples[row][col] > 0 {
				f[row][col] = 1
			}
		}
	}

	for remain := 1; remain < k; remain++ {
		g := make([][]int, rows)
		for i, _ := range g {
			g[i] = make([]int, cols)
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				for nextRow := row + 1; nextRow < rows; nextRow++ {
					if apples[row][col] - apples[nextRow][col] > 0 {
						g[row][col] += f[nextRow][col]
						g[row][col] %= modulo
					}
				}

				for nextCol := col + 1; nextCol < cols; nextCol++ {
					if apples[row][col] - apples[row][nextCol] > 0 {
						g[row][col] += f[row][nextCol]
						g[row][col] %= modulo
					}
				}
			}
		}

		for i, _ := range f {
			copy(f[i], g[i])
		}
	}

	return f[0][0]
}

func main() {
	// result: 3
	// pizza := []string{"A..","AAA","..."}
	// k := int(3)

	// result: 1
	// pizza := []string{"A..","AA.","..."}
	// k := int(3)

	// result: 1
	pizza := []string{"A..","A..","..."}
	k := int(1)

	// result: 
	// pizza := []string{}
	// k := int(0)

	result := ways(pizza, k)
	fmt.Printf("result = %v\n", result)
}

