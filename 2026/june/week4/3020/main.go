package main

import (
	"fmt"
)

func maximumLength(nums []int) int {
	freq := make(map[int]int)
	distinct := make([]int, 0, len(nums))
	for _, num := range nums {
		if freq[num] == 0 {
			distinct = append(distinct, num)
		}
		freq[num]++
	}

	result := 1
	if freq[1] > 0 {
		candidate := freq[1]
		if candidate % 2 == 0 {
			candidate--
		}
		if candidate > result {
			result = candidate
		}
	}

	for _, start := range distinct {
		if start == 1 {
			continue
		}

		length := int(1)
		cur := int64(start)
		for {
			next := cur * cur
			if next > 1_000_000_000 {
				break
			}

			if freq[int(cur)] < 2 {
				break
			}

			if freq[int(next)] == 0 {
				break
			}

			length++
			cur = next
		}

		candidate := 2 * length - 1
		if candidate > result {
			result = candidate
		}
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{5,4,1,2,2}

	// result: 1
	// nums := []int{1,3,2,4}

	// result: 2
	nums := []int{1,1}

	// result: 
	// nums := []int{}

	result := maximumLength(nums)
	fmt.Printf("result = %v\n", result)
}
