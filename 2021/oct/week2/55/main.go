package main
import (
	"fmt"
)

func canJump(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// fmt.Printf("nums[%v] = %v\n", i, num)
		if i + num >= len(nums) - 1 {
			break
		}
		for j := i + num; j > i; j-- {
			// fmt.Printf("nums[%v] = %v\n", j, nums[j])
			if nums[j] > 0 {
				if canJump(nums[j:]) {
					return true
				}
			}
		}
		return false
	}
    return true
}

func main() {
	// result: true
	// nums := []int{2, 3, 1, 1, 4}

	// result: false
	// nums := []int{3, 2, 1, 0, 4}

	// result: true
	// nums := []int{2, 2, 0, 1}

	// result: true
	nums := []int{2,2,0,1,3,13,8,1,7,4,0,3,2,1,1,0,0,2,3}

	// result: 
	// nums := []int{}

	result := canJump(nums)
	fmt.Printf("result = %v\n", result)
}

