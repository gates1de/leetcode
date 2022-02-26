package main
import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)

	pre := nums[0]
	result := pre
	maxCount := int(1)
	count := int(1)
	for _, num := range nums[1:] {
		if pre == num {
			count++
			continue
		}

		if maxCount < count {
			maxCount = count
			result = pre
		}
		pre = num
		count = 1
	}

	if maxCount < count {
		result = nums[len(nums) - 1]
	}
	// fmt.Println(count, maxCount)

	return result
}

func main() {
	// result: 3
	// nums := []int{3, 2, 3}

	// result: 2
	// nums := []int{2, 2, 1, 1, 1, 2, 2}

	// result: 1
	// nums := []int{1}

	// result: 3
	nums := []int{3,9,3,2,1,3,7,6,3,4,3,8,2,3,3,4,3,9,8,1,3,3,3,3,3}

	// result: 
	// nums := []int{}

	result := majorityElement(nums)
	fmt.Printf("result = %v\n", result)
}

