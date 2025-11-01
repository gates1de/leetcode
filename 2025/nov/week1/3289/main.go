package main

import (
	"fmt"
)

func getSneakyNumbers(nums []int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	result := make([]int, 0, len(m))
	for num, count := range m {
		if count == 2 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// result: [0,1]
	// nums := []int{0,1,1,0}


	// result: [2,3]
	// nums := []int{0,3,2,1,3,2}

	// result: [4,5]
	nums := []int{7,1,5,4,3,4,6,0,9,5,8,2}

	// result: []
	// nums := []int{}

	result := getSneakyNumbers(nums)
	fmt.Printf("result = %v\n", result)
}

