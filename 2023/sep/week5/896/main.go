package main
import (
	"fmt"
)

func isMonotonic(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	pre := nums[0]
	// 0 = init, 1 = ascending, 2 = descending
	monotonicType := int(0)

	for _, num := range nums[1:] {
		if pre == num {
			continue
		}

		if monotonicType == 0 {
			if pre < num {
				monotonicType = 1
			} else {
				monotonicType = 2
			}
		} else if monotonicType == 1 {
			if pre > num {
				return false
			}
		} else {
			if pre < num {
				return false
			}
		}

		pre = num
	}

	return true
}

func main() {
	// result: true
	// nums := []int{1,2,2,3}

	// result: true
	// nums := []int{6,5,4,4}

	// result: false
	nums := []int{1,3,2}

	// result: 
	// nums := []int{}

	result := isMonotonic(nums)
	fmt.Printf("result = %v\n", result)
}

