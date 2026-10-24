package main

import (
	"fmt"
)

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	answer := -1
	k := (n - nums[0]) % n
	sorted := true
	for i := range nums {
		if nums[(i+k)%n] != i {
			sorted = false
			break
		}
	}
	if sorted {
		answer = k
		if candidate := n - k + 2; candidate < answer {
			answer = candidate
		}
	}

	zeroIndex := 0
	for i, value := range nums {
		if value == 0 {
			zeroIndex = i
			break
		}
	}
	k = (n - 1 - zeroIndex + n) % n
	reversed := true
	for i := range nums {
		if nums[n-1-(i+k)%n] != i {
			reversed = false
			break
		}
	}
	if reversed {
		candidate := k + 1
		if n-k+1 < candidate {
			candidate = n - k + 1
		}
		if answer == -1 || candidate < answer {
			answer = candidate
		}
	}

	return answer
}

func main() {
	// result: 2
	// nums := []int{0,2,1}

	// result: 2
	// nums := []int{1,0,2}

	// result: -1
	nums := []int{2, 0, 1, 3}

	// result: 0
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}
