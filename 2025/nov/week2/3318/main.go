package main

import (
	"fmt"
)

func findXSum(nums []int, k int, x int) []int {
	result := make([]int, 0, len(nums))
	for i := 0; i <= len(nums) - k; i++ {
		result = append(result, xsum(nums[i:i + k], x))
	}

	return result
}

func xsum(nums []int, x int) int {
	counter := make([]int, 51)
	sum := int(0)

	for _, n := range nums {
		counter[n]++
	}

	for x != 0 {
		count, num := max(counter[:])
		sum += count * num
		counter[num] = 0
		x--
	}

	return sum
}

func max(nums []int) (int, int) {
	num, freq := -1, -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > freq {
			freq = nums[i]
			num = i
		}
	}

	return freq, num
}

func main() {
	// result: [6,10,12]
	// nums := []int{1,1,2,2,3,4,2,3}
	// k := int(6)
	// x := int(2)

	// result: [11,15,15,15,12]
	nums := []int{3,8,7,8,7,5}
	k := int(2)
	x := int(2)

	// result: []
	// nums := []int{}
	// k := int(0)
	// x := int(0)

	result := findXSum(nums, k, x)
	fmt.Printf("result = %v\n", result)
}

