package main
import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := map[int]int{}
	helper(0, m, nums)
	max := math.MinInt32
	for _, val := range m {
		if max < val {
			max = val
		}
	}
	return max
}

func helper(current int, m map[int]int, nums []int) {
	if len(nums) == 0 {
		return
	}
	num := nums[0]

	if current == 0 {
		if m[len(nums)] <= num {
			m[len(nums)] = num
		}
		helper(num, m, nums[1:])
		return
	}

	if num <= 0 {
		helper(0, m, nums[1:])
	}

	if m[len(nums)] <= current * num {
		m[len(nums)] = current * num
	}

	helper(current * num, m, nums[1:])
}

func main() {
	// result: 6
	// nums := []int{2, 3, -2, 4}

	// result: 0
	// nums := []int{-2, 0, -1}

	// result: 6
	// nums := []int{3, 2, -1, -1, 0, 4, 1}

	// result: 8
	// nums := []int{3, 2, -1, 1, 8, 1}

	// result: 1
	// nums := []int{-1, -1, -1}

	// result: 24
	// nums := []int{2, -5, -2, -4, 3}

	// result: -1
	// nums := []int{-2}

	// result: 60
	nums := []int{1, 0, -1, 2, 3, -5, -2}

	// result: 
	// nums := []int{}

	result := maxProduct(nums)
	fmt.Printf("result = %v\n", result)
}

