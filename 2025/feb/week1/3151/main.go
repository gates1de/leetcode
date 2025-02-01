package main
import (
	"fmt"
)

func isArraySpecial(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}

	pre := nums[0]
	for _, num := range nums[1:] {
		if (pre % 2 == 0 && num % 2 == 1) || (pre % 2 == 1 && num % 2 == 0) {
			pre = num
			continue
		}

		return false
	}

	return true
}

func main() {
	// result: true
	// nums := []int{1}

	// result: true
	// nums := []int{2,1,4}

	// result: false
	nums := []int{4,3,1,6}

	// result: 
	// nums := []int{}

	result := isArraySpecial(nums)
	fmt.Printf("result = %v\n", result)
}

