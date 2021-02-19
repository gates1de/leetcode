package main
import (
	"fmt"
	"sort"
)

func minimumDeviation(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	min := nums[0]
	max := nums[len(nums) - 1]
	for i, num := range nums {
		if num % 2 == 0 && num <= max {
			for num % 2 == 0 && min <= num / 2 {
				num /= 2
			}
			nums[i] = num
		}
	}
	sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	min = nums[0]
	max = nums[len(nums) - 1]
	result := abs(max, min)
	if min % 2 == 1 && result >= abs(max, 2 * min) {
		result = abs(max, 2 * min)
	}
	return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func abs(a int, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{4, 1, 5, 20, 3}
	// nums := []int{2, 10, 8}
    // nums := []int{3, 5}
	// nums := []int{4, 9, 4, 5}
	nums := []int{10, 4, 3}

	result := minimumDeviation(nums)
	fmt.Printf("result = %v\n", result)
}

