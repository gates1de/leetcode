package main
import (
	"fmt"
)

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	if m == 0 {
		return img
	}

	n := len(img[0])
	if n == 0 {
		return img
	}


	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
	}

	for i, _ := range img {
		for j, _ := range img[i] {
			sum := int(0)
			count := int(0)

			for x := i - 1; x <= i + 1; x++ {
				for y := j - 1; y <= j + 1; y++ {
					if 0 <= x && x < m && 0 <= y && y < n {
						sum += img[x][y]
						count += 1
					}
				}
			}

			result[i][j] = sum / count
		}
	}

	return result
}


func main() {
	// result: [[0,0,0],[0,0,0],[0,0,0]]
	// img := [][]int{{1,1,1},{1,0,1},{1,1,1}}

	// result: [[137,141,137],[141,138,141],[137,141,137]]
	img := [][]int{{100,200,100},{200,50,200},{100,200,100}}

	// result: 
	// img := [][]int{}

	result := imageSmoother(img)
	fmt.Printf("result = %v\n", result)
}

