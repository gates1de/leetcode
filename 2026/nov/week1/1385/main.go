package main

import (
	"fmt"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	result := 0
	for _, value := range arr1 {
		left := sort.SearchInts(arr2, value-d)
		if left == len(arr2) || arr2[left] > value+d {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// arr1 := []int{4,5,8}
	// arr2 := []int{10,9,1,8}
	// d := 2

	// result: 2
	// arr1 := []int{1,4,2,3}
	// arr2 := []int{-4,-3,6,10,20,30}
	// d := 3

	// result: 1
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	d := 6

	// result: 0
	// arr1 := []int{}
	// arr2 := []int{}
	// d := 0

	result := findTheDistanceValue(arr1, arr2, d)
	fmt.Printf("result = %v\n", result)
}
