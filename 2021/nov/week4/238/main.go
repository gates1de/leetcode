package main
import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
    left := make([]int, len(nums))
    right := make([]int, len(nums))
    left[0] = 1
    right[len(right) - 1] = 1

    for i := 1; i< len(left); i++ {
        left[i] = left[i - 1] * nums[i - 1]
    }
    for i := len(right) - 2; i >= 0; i-- {
        right[i] = right[i + 1] * nums[i + 1]
    }
    for i := 0; i < len(nums); i++ {
        nums[i] = left[i] * right[i]
    }
    return nums
}

// Accepted, but using the division operation
func mySolution(nums []int) []int {
	allProduct := int(1)
	isExistsZero := false
	result := make([]int, len(nums))
	for _, num := range nums {
		if num == 0 {
			if isExistsZero {
				return result
			}
			isExistsZero = true
			continue
		}
		allProduct *= num
	}

	for i, num := range nums {
		if num == 0 {
			result[i] = allProduct
			continue
		}
		if isExistsZero {
			result[i] = 0
			continue
		}
		result[i] = allProduct / num
	}

	return result
}

func main() {
	// result: [24,12,8,6]
	// nums := []int{1, 2, 3, 4}

	// result: [0,0,9,0,0]
	// nums := []int{-1, 1, 0, -3, 3}

	// result: [-2,2,-1]
	// nums := []int{-1, 1, -2}

	// result: [0,0,0,0,0]
	nums := []int{0, 0, -3, -2, 2}

	// result: 
	// nums := []int{}

	result := productExceptSelf(nums)
	fmt.Printf("result = %v\n", result)
}

