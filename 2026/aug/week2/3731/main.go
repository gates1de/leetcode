package main

import (
	"fmt"
)

const maximumInputValue = int(100)

func findMissingElements(nums []int) []int {
	present := [maximumInputValue + 1]bool{}
	minimum, maximum := nums[0], nums[0]
	for _, num := range nums {
		present[num] = true
		if num < minimum {
			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}

	result := make([]int, 0, maximum - minimum + 1 - len(nums))
	for offset := range maximum - minimum + 1 {
		num := minimum + offset
		if !present[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// result: [3]
	// nums := []int{1,4,2,5}

	// result: []
	// nums := []int{7,8,6,9}

	// result: [2,3,4]
	nums := []int{5, 1}

	// result: []
	// nums := []int{}

	result := findMissingElements(nums)
	fmt.Printf("result = %v\n", result)
}
