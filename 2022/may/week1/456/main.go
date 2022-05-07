package main
import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	max := math.MinInt32
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			return true
		}
		for len(stack) > 0 && stack[len(stack) - 1] < nums[i] {
			max = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

// Time Limit Exceeded
func ngSolution(nums []int) bool {
	n := len(nums)
	if n <= 2 {
		return false
	}

	m := map[int]int{}
	min := nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		if num < min {
			min = num
		} else {
			m[i] = min
		}
	}

	min = nums[n - 1]
	max := min
	for i := n - 2; i >= 0; i-- {
		num := nums[i]
		if num < min {
			min = num
		} else {
			if val, ok := m[i]; ok && val < num {
				for j := i + 1; j < n; j++ {
					if val < nums[j] && num > nums[j] {
						return true
					}
				}
			}

			if num > max {
				max = num
			}
		}
	}

	return false
}

func main() {
	// result: false
	// nums := []int{1, 2, 3, 4}

	// result: true
	// nums := []int{3, 1, 4, 2}

	// result: true
	// nums := []int{-1, 3, 2, 0}

	// result: false
	// nums := []int{1, 0, 1, -4, -3}

	// result: true
	// nums := []int{3, 5, 0, 3, 4}

	// result: false
	// nums := []int{-2, 1, 1}

	// result: true
	// nums := []int{1,3,2,4,5,6,7,8,9,10}

	// result: true
	nums := []int{-2, 1, 2, -2, 1, 2}

	// result: 
	// nums := []int{}

	result := find132pattern(nums)
	fmt.Printf("result = %v\n", result)
}

