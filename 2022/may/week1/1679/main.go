package main
import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i := int(0)
	j := len(nums) - 1
	result := int(0)

	for i < j {
		left := nums[i]
		right := nums[j]
		if left >= k || right * 2 < k {
			break
		}

		if left + right == k {
			result++
			i++
			j--
		} else if left + right < k {
			i++
		} else {
			j--
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{1, 2, 3, 4}
	// k := int(5)

	// result: 1
	// nums := []int{3, 1, 3, 4, 3}
	// k := int(6)

	// result: 0
	// nums := []int{3,1,5,1,1,1,1,1,2,2,3,2,2}
	// k := int(1)

	// result: 2
	nums := []int{2, 1, 5, 4, 6, 7}
	k := int(6)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

