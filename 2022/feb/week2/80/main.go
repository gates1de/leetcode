package main
import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count := int(1)
	result := int(1)
	for i := 1; i < len(nums); i++ {
		pre := nums[i - 1]
		num := nums[i]

		if pre == num {
			if count == 2 {
				nums = remove(nums, i)
				i--
				continue
			}
			count++
			result++
		} else {
			count = 1
			result++
		}
	}

	// fmt.Println(nums)
	return result
}

func remove(slice []int, i int) []int {
    return append(slice[:i], slice[i + 1:]...)
}

func main() {
	// result: 5
	// nums := []int{1, 1, 1, 2, 2, 3}

	// result: 7
	// nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	// result: 1
	// nums := []int{0}

	// result: 2
	// nums := []int{1, 2}

	// result: 4
	nums := []int{1, 1, 3, 3, 3}

	// result: 
	// nums := []int{}

	result := removeDuplicates(nums)
	fmt.Printf("result = %v\n", result)
}

