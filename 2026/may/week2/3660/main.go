package main

import (
	"fmt"
)

type Item struct {
	Value int
	Left  int
	Right int
}

func maxValue(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	stack := make([]Item, 0)
	for i := range n {
		curr := Item{Value: nums[i], Left: i, Right: i}

		for len(stack) > 0 && stack[len(stack) - 1].Value > nums[i] {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			if top.Value > curr.Value {
				curr.Value = top.Value
			}
			curr.Left = top.Left
		}

		stack = append(stack, curr)
	}

	for i := range len(stack) {
		for j := stack[i].Left; j <= stack[i].Right; j++ {
			result[j] = stack[i].Value
		}
	}

	return result
}

func main() {
	// result: [2,2,3]
	// nums := []int{2,1,3}

	// result: [3,3,3]
	nums := []int{2,3,1}

	// result: []
	// nums := []int{}

	result := maxValue(nums)
	fmt.Printf("result = %v\n", result)
}

