package main
import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	nums := make([]int, 1001)
	result := make([]int, 1001)
	for _, num := range nums1 {
		nums[num]++
	}

	count := int(0)
	for _, num := range nums2 {
		if nums[num] > 0 {
			result[count] = num
			count++
			nums[num]--
		}
	}

	return result[:count]
}

func main() {
	// result: [2,2]
	// nums1 := []int{1,2,2,1}
	// nums2 := []int{2,2}

	// result: [4,9]
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	// result: []
	// nums1 := []int{}
	// nums2 := []int{}

	result := intersect(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

