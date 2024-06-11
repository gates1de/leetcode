package main
import (
	"fmt"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	maxVal := int(0)
	for _, num := range arr1 {
		maxVal = max(maxVal, num)
	}

	counter := make([]int, maxVal + 1)
	for _, num := range arr1 {
		counter[num]++
	}

	result := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		for counter[num] > 0 {
			result = append(result, num)
			counter[num]--
		}
	}

	for num := 0; num <= maxVal; num++ {
		for counter[num] > 0 {
			result = append(result, num)
			counter[num]--
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [2,2,2,1,4,3,3,9,6,7,19]
	// arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	// arr2 := []int{2,1,4,3,9,6}

	// result: [22,28,8,6,17,44]
	arr1 := []int{28,6,22,8,44,17}
	arr2 := []int{22,28,8,6}

	// result: []
	// arr1 := []int{}
	// arr2 := []int{}

	result := relativeSortArray(arr1, arr2)
	fmt.Printf("result = %v\n", result)
}

