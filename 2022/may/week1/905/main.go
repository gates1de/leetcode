package main
import (
	"fmt"
)

func sortArrayByParity(nums []int) []int {
	i := int(0)
	j := len(nums) - 1
	for i < j {
		left := nums[i]
		right := nums[j]
		if left % 2 == 0 {
			i++
		} else {
			if right % 2 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			} else {
				j--
			}
		}
	}

	return nums
}

func main() {
	// result: [2,4,3,1] ([4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted)
	// nums := []int{3, 1, 2, 4}

	// result: [0]
	// nums := []int{0}

	// result: [2,4,6,3]
	// nums := []int{2, 4, 6, 3}

	// result: [6,1,3,5]
	// nums := []int{1, 3, 5, 6}

	// result: [2,4,3,5,1]
	nums := []int{5, 4, 3, 2, 1}

	// result: 
	// nums := []int{}

	result := sortArrayByParity(nums)
	fmt.Printf("result = %v\n", result)
}

