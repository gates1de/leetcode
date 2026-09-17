package main

import (
	"fmt"
)

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	inf := n + 1
	bestPrefix := make([]int, n)
	best := inf
	answer := inf
	left := 0
	sum := 0

	for right, value := range arr {
		sum += value
		for sum > target {
			sum -= arr[left]
			left++
		}

		if sum == target {
			length := right - left + 1
			if left > 0 && bestPrefix[left-1] != inf {
				candidate := bestPrefix[left-1] + length
				if candidate < answer {
					answer = candidate
				}
			}
			if length < best {
				best = length
			}
		}
		bestPrefix[right] = best
	}

	if answer == inf {
		return -1
	}
	return answer
}

func main() {
	// result: 2
	// arr := []int{3,2,2,4,3}
	// target := int(3)

	// result: 2
	// arr := []int{7,3,4,7}
	// target := int(7)

	// result: -1
	arr := []int{4, 3, 2, 6, 2, 3, 4}
	target := int(6)

	// result: 0
	// arr := []int{}
	// target := int(0)

	result := minSumOfLengths(arr, target)
	fmt.Printf("result = %v\n", result)
}
