package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)
	sort.Ints(arr)
	minNum := math.MaxInt32
	result := make([][]int, 0, n)

	for i := range len(arr) - 1 {
		diff := arr[i + 1] - arr[i]
		if diff < minNum {
			minNum = diff
			result = make([][]int, 0, n)
			result = append(result, []int{arr[i], arr[i + 1]})
		} else if diff == minNum {
			result = append(result, []int{arr[i], arr[i + 1]})
		}
	}

	return result
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

