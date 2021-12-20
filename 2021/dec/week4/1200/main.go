package main
import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	m := map[int][][]int{}
	min := math.MaxInt32
	for i := 0; i < len(arr) - 1; i++ {
		n1 := arr[i]
		n2 := arr[i + 1]
		a := abs(n1, n2)
		if m[a] == nil {
			m[a] = [][]int{}
		}
		m[a] = append(m[a], []int{n1, n2})
		if min > a {
			min = a
		}
	}

	return m[min]
}

func abs(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	// result: [[1,2],[2,3],[3,4]]
	// arr := []int{4, 2, 1, 3}

	// result: [[1,3]]
	// arr := []int{1, 3, 6, 10, 15}

	// result: [[-14,-10],[19,23],[23,27]]
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}

	// result: 
	// arr := []int{}

	result := minimumAbsDifference(arr)
	fmt.Printf("result = %v\n", result)
}

