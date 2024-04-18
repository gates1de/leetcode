package main
import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
    ver := map[[2]int]bool{}
    hor := map[[2]int]bool{}
    result := int(0)
    for i, row := range grid {
        for j, cell := range row {
            if cell == 0 {
                continue
            }

            if !ver[[2]int{i, j}] {
                result++
                ver[[2]int{i, j}] = true
            } else {
                result--
            }
            if !ver[[2]int{i + 1, j}] {
                result++
                ver[[2]int{i + 1, j}] = true
            } else {
                result--
            }

            if !hor[[2]int{i, j}] {
                result++
                hor[[2]int{i, j}] = true
            } else {
                result--
            }
            if !hor[[2]int{i, j + 1}] {
                result++
                hor[[2]int{i, j + 1}] = true
            } else {
                result--
            }
        }
    }

    return result
}

func main() {
	// result: 16
	// grid := [][]int{
	// 	{0, 1, 0, 0},
	// 	{1, 1, 1, 0},
	// 	{0, 1, 0, 0},
	// 	{1, 1, 0, 0},
	// }

	// result: 4
	// grid := [][]int{{1}}

	// result: 4
	grid := [][]int{{1,0}}

	// result: 
	// grid := [][]int{
	// }

	result := islandPerimeter(grid)
	fmt.Printf("result = %v\n", result)
}

