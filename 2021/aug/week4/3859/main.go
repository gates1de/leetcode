package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func rectangleArea(rectangles [][]int) int {
	if len(rectangles) == 1 {
		rect := rectangles[0]
		return (rect[2] - rect[0]) * (rect[3] - rect[1]) % modulo
	}

	m := map[[2]int]bool{}
	for _, rect := range rectangles {
		for i := rect[0]; i < rect[2]; i++ {
			for j:= rect[1]; j < rect[3]; j++ {
				m[[2]int{i, j}] = true
			}
		}
	}
	result := int(0)
	for range m {
		result++
	}
	return result % modulo
}

func main() {
	// result: 6
	// rectangles := [][]int{
	// 	{0, 0, 2, 2},
	// 	{1, 0, 2, 3},
	// 	{1, 0, 3, 1},
	// }

	// result: 49
	rectangles := [][]int{
		{0, 0, 1000000000, 1000000000},
	}

	// result: 
	// rectangles := [][]int{
	// 	{, , , },
	// 	{, , , },
	// 	{, , , },
	// }

	result := rectangleArea(rectangles)
	fmt.Printf("result = %v\n", result)
}

