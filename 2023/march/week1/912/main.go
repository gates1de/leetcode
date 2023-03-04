package main
import (
	"fmt"
)

func sortArray(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	nums1 := make([]int, len(nums[:n / 2]))
	copy(nums1, nums[:n / 2])
	nums1 = sortArray(nums1)
	nums2 := make([]int, len(nums[n / 2:]))
	copy(nums2, nums[n / 2:])
	nums2 = sortArray(nums2)

	i := int(0)
	j := int(0)
	k := int(0)
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			nums[k] = nums2[j]
			j++
			k++
			continue
		} else if j >= len(nums2) {
			nums[k] = nums1[i]
			i++
			k++
			continue
		}

		num1 := nums1[i]
		num2 := nums2[j]
		if num1 <= num2 {
			nums[k] = num1
			i++
		} else {
			nums[k] = num2
			j++
		}
		k++
	}

	return nums
}

func main() {
	// result: [1,2,3,5]
	// nums := []int{5,2,3,1}

	// result: [0,0,1,1,2,5]
	nums := []int{5,1,1,2,0,0}

	// result: 
	// nums := []int{}

	result := sortArray(nums)
	fmt.Printf("result = %v\n", result)
}

