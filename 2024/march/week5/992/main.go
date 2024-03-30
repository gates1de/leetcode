package main
import (
	"fmt"
)

func subarraysWithKDistinct(nums []int, k int) int {
	return helper(nums, k) - helper(nums, k - 1)
}

func helper(nums []int, k int) int {
	frequencies := make(map[int]int)
	left := int(0)
	result := int(0)

	for right := 0; right < len(nums); right++ {
		frequencies[nums[right]]++

		for len(frequencies) > k {
			frequencies[nums[left]]--
			if frequencies[nums[left]] == 0 {
				delete(frequencies, nums[left])
			}
			left++
		}

		result += right - left + 1
	}

	return result
}

func main() {
	// result: 7
	// nums := []int{1,2,1,2,3}
	// k := int(2)

	// result: 3
	nums := []int{1,2,1,3,4}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := subarraysWithKDistinct(nums, k)
	fmt.Printf("result = %v\n", result)
}

