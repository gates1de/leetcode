package main

import (
	"fmt"
)

func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	nLayer := min(m / 2, n / 2)

	for layer := range nLayer {
			r := make([]int, 0)
			c := make([]int, 0)
			val := make([]int, 0)

			for i := layer; i < m - layer - 1; i++ {
					r = append(r, i)
					c = append(c, layer)
					val = append(val, grid[i][layer])
			}
			for j := layer; j < n - layer - 1; j++ {
					r = append(r, m - layer - 1)
					c = append(c, j)
					val = append(val, grid[m-layer - 1][j])
			}
			for i := m - layer - 1; i > layer; i-- {
					r = append(r, i)
					c = append(c, n - layer - 1)
					val = append(val, grid[i][n - layer - 1])
			}
			for j := n - layer - 1; j > layer; j-- {
					r = append(r, layer)
					c = append(c, j)
					val = append(val, grid[layer][j])
			}

			total := len(val)
			tmp := k % total
			for i := range total {
					index := (i + total - tmp) % total
					grid[r[i]][c[i]] = val[index]
			}
	}

	return grid
}

func main() {
  // result: [[10,20],[40,30]]
  // grid := [][]int{{40,10},{30,20}}
  // k := int(1)

  // result: [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
  grid := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}
  k := int(2)

  // result: []
  // grid := [][]int{}
  // k := int(0)

	result := rotateGrid(grid, k)
	fmt.Printf("result = %v\n", result)
}

