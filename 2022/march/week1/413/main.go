package main
import (
	"fmt"
)

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	result := int(0)
	count := int(2)
	preDiff := nums[1] - nums[0]
	pre := nums[1]
	for _, num := range nums[2:] {
		diff := num - pre
		pre = num
		if diff != preDiff {
			if count >= 3 {
				result += calc(count)
			}
			count = 2
			preDiff = diff
			continue
		}

		count++
	}

	if count >= 3 {
		result += calc(count)
	}

	return result
}

func calc(n int) int {
	result := int(0)
	for i := n; i >= 3; i-- {
		result += n - i + 1
	}
	return result
}

func main() {
	// result: 3
	// nums := []int{1, 2, 3, 4}

	// result: 0
	// nums := []int{1}

	// result: 0
	// nums := []int{1, 2}

	// result: 0
	// nums := []int{1, 2, 2}

	// result: 1
	// nums := []int{1, 2, 3}

	// result: 
	nums := []int{1, 2, 3, 4, -1, -1, -1, -1, -2, -3, -4, -5}

	// result: 
	// nums := []int{}

	result := numberOfArithmeticSlices(nums)
	fmt.Printf("result = %v\n", result)
}

