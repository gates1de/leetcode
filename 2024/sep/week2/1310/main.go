package main
import (
	"fmt"
)

func xorQueries(arr []int, queries [][]int) []int {
	result := make([]int, 0, len(arr))
	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i - 1]
	}

	for _, q := range queries {
		if q[0] > 0 {
			result = append(result, arr[q[0] - 1] ^ arr[q[1]])
		} else {
			result = append(result, arr[q[1]])
		}
	}

	return result
}

func main() {
	// result: [2,7,14,8]
	// arr := []int{1,3,4,8}
	// queries := [][]int{{0,1},{1,2},{0,3},{3,3}}

	// result: [8,0,4,4]
	arr := []int{4,8,2,10}
	queries := [][]int{{2,3},{1,3},{0,0},{0,3}}

	// result: []
	// arr := []int{}
	// queries := [][]int{}

	result := xorQueries(arr, queries)
	fmt.Printf("result = %v\n", result)
}

