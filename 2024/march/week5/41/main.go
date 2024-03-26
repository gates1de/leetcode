package main
import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		num := nums[i]
		correctIndex := num - 1
		if num > 0 && num <= n && num != nums[correctIndex] {
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
			i--
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	// result: 3
	// nums := []int{1,2,0}

	// result: 2
	// nums := []int{3,4,-1,1}

	// result: 1
	nums := []int{7,8,9,11,12}

	// result: 
	// nums := []int{}

	result := firstMissingPositive(nums)
	fmt.Printf("result = %v\n", result)
}

