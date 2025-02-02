package main
import (
	"fmt"
)

func check(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	inversionCount := int(0)

	for i := 1; i < n; i++ {
		if nums[i] < nums[i - 1] {
			inversionCount++
		}
	}

	if nums[0] < nums[n - 1] {
		inversionCount++
	}

	return inversionCount <= 1
}

func main() {
	// result: true
	// nums := []int{3,4,5,1,2}

	// result: false
	// nums := []int{2,1,3,4}

	// result: true
	nums := []int{1,2,3}

	// result: 
	// nums := []int{}

	result := check(nums)
	fmt.Printf("result = %v\n", result)
}

