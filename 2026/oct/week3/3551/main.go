package main

import (
	"fmt"
	"sort"
)

func minSwaps(nums []int) int {
	sorted := append([]int(nil), nums...)
	sort.Slice(sorted, func(i, j int) bool {
		sumI := digitSum(sorted[i])
		sumJ := digitSum(sorted[j])
		if sumI != sumJ {
			return sumI < sumJ
		}
		return sorted[i] < sorted[j]
	})

	position := make(map[int]int, len(nums))
	for i, value := range sorted {
		position[value] = i
	}

	visited := make([]bool, len(nums))
	result := 0
	for i := range nums {
		if visited[i] {
			continue
		}
		cycleLength := 0
		current := i
		for !visited[current] {
			visited[current] = true
			current = position[nums[current]]
			cycleLength++
		}
		result += cycleLength - 1
	}

	return result
}

func digitSum(value int) int {
	result := 0
	for value > 0 {
		result += value % 10
		value /= 10
	}
	return result
}

func main() {
	// result: 1
	// nums := []int{37,100}

	// result: 0
	// nums := []int{22,14,33,7}

	// result: 2
	nums := []int{18, 43, 34, 16}

	// result: 0
	// nums := []int{}

	result := minSwaps(nums)
	fmt.Printf("result = %v\n", result)
}
