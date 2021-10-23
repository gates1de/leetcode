package main
import (
	"fmt"
)

func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1
	result := int(5000)
    for l < r {
        m := (r + l) / 2
		// fmt.Printf("l = %v, r = %v\n", l, r)
        if nums[l] < nums[r] && result > nums[l] {
			result = nums[l]
        }

        if m > 0 && nums[m] < nums[m - 1] && result > nums[m] {
			result = nums[m]
        }

		if nums[r] == nums[m] {
			l++
		} else if nums[r] < nums[m] {
            l = m + 1
        } else {
            r = m
        }
    }

	if result > nums[l] {
		result = nums[l]
	}
    return result
}

func main() {
	// result: 1
	// nums := []int{1, 3, 5}

	// result: 0
	// nums := []int{2, 2, 2, 0, 1}

	// result: 0
    // nums := []int{4, 5, 6, 7, 0, 1, 2}

	// result: -3300
    // nums := []int{400,500,600,700,800,900,1000,1100,2200,-3300,-440,-85,-78,-77}

	// result: 1
	// nums := []int{3, 3, 1, 3}

	// result: 1
	nums := []int{10, 1, 10, 10, 10}

	// result: 
	// nums := []int{}

	result := findMin(nums)
	fmt.Printf("result = %v\n", result)
}

