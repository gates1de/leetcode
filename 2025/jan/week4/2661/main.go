package main
import (
	"fmt"
	"math"
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	numToIndex := make(map[int]int)
	for i, num := range arr {
		numToIndex[num] = i
	}

	result := math.MaxInt32
	m := len(mat)
	n := len(mat[0])

	for row := 0; row < m; row++ {
		lastElementIndex := math.MinInt32

		for col := 0; col < n; col++ {
			index := numToIndex[mat[row][col]]
			lastElementIndex = max(lastElementIndex, index)
		}

		result = min(result, lastElementIndex)
	}

	for col := 0; col < n; col++ {
		lastElementIndex := math.MinInt32

		for row := 0; row < m; row++ {
			index := numToIndex[mat[row][col]]
			lastElementIndex = max(lastElementIndex, index)
		}

		result = min(result, lastElementIndex)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// arr := []int{1,3,4,2}
	// mat := [][]int{{1,4},{2,3}}

	// result: 3
	arr := []int{2,8,7,4,1,3,5,6,9}
	mat := [][]int{{3,2,5},{1,4,6},{8,7,9}}

	// result: 
	// arr := []int{}
	// mat := [][]int{}

	result := firstCompleteIndex(arr, mat)
	fmt.Printf("result = %v\n", result)
}

