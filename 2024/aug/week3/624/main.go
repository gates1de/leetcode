package main
import (
	"fmt"
)

func maxDistance(arrays [][]int) int {
	m := len(arrays)
	if m < 2 || len(arrays[0]) == 0 {
		return -1
	}

	smallest := arrays[0][0]
	biggest := arrays[0][len(arrays[0]) - 1]
	result := int(0)

	for i := 1; i < m; i++ {
		result = max(result, abs(arrays[i][len(arrays[i]) - 1] - smallest))
		result = max(result, abs(biggest - arrays[i][0]))
		smallest = min(smallest, arrays[i][0])
		biggest = max(biggest, arrays[i][len(arrays[i]) - 1])
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

func abs(num int) int {
	if num <= 0 {
		return -num
	}
	return num
}

func main() {
	// result: 4
	// arrays := [][]int{{1,2,3},{4,5},{1,2,3}}

	// result: 0
	arrays := [][]int{{1},{1}}

	// result: 
	// arrays := [][]int{}

	result := maxDistance(arrays)
	fmt.Printf("result = %v\n", result)
}

