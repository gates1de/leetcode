package main

import (
	"fmt"
)

type Answer struct {
	Coordinate [3]int
}

func (this *Answer) Put(x int) {
	if x > this.Coordinate[0] {
		this.Coordinate[2] = this.Coordinate[1]
		this.Coordinate[1] = this.Coordinate[0]
		this.Coordinate[0] = x
	} else if x != this.Coordinate[0] && x > this.Coordinate[1] {
		this.Coordinate[2] = this.Coordinate[1]
		this.Coordinate[1] = x
	} else if x != this.Coordinate[0] && x != this.Coordinate[1] && x > this.Coordinate[2] {
		this.Coordinate[2] = x
	}
}

func (this *Answer) Get() []int {
	result := make([]int, 0, len(this.Coordinate))
	for _, num := range this.Coordinate {
		if num != 0 {
			result = append(result, num)
		}
	}
	return result
}

func getBiggestThree(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	sum1 := make([][]int, m + 1)
	sum2 := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		sum1[i] = make([]int, n + 2)
		sum2[i] = make([]int, n + 2)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum1[i][j] = sum1[i - 1][j - 1] + grid[i - 1][j - 1]
			sum2[i][j] = sum2[i - 1][j + 1] + grid[i - 1][j - 1]
		}
	}

	result := Answer{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result.Put(grid[i][j])

			for k := i + 2; k < m; k += 2 {
				ux, uy := i, j
				dx, dy := k, j
				lx, ly := (i + k) / 2, j - (k - i) / 2
				rx, ry := (i + k) / 2, j + (k - i) / 2
				if ly < 0 || ry >= n {
					break
				}

				sum := (sum2[lx + 1][ly + 1] - sum2[ux][uy + 2]) +
					(sum1[rx + 1][ry + 1] - sum1[ux][uy]) +
					(sum1[dx + 1][dy + 1] - sum1[lx][ly]) +
					(sum2[dx + 1][dy + 1] - sum2[rx][ry + 2]) -
					(grid[ux][uy] + grid[dx][dy] + grid[lx][ly] + grid[rx][ry])

				result.Put(sum)
			}
		}
	}

	return result.Get()
}

func main() {
	// result: [228,216,211]
	// grid := [][]int{{3,4,5,1,3},{3,3,4,2,3},{20,30,200,40,10},{1,5,5,4,1},{4,3,2,2,5}}

	// result: [20,9,8]
	// grid := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: [7]
	grid := [][]int{{7,7,7}}

	// result: []
	// grid := [][]int{}

	result := getBiggestThree(grid)
	fmt.Printf("result = %v\n", result)
}

