package main
import (
	"fmt"
)

func minSum(nums1 []int, nums2 []int) int64 {
	sum := int64(0)
	result := int64(0)
	zero1 := int(0)
	zero2 := int(0)

	for _, num := range nums1 {
		sum += int64(num)
		if num == 0 {
			sum++
			zero1++
		}
	}

	for _, num := range nums2 {
		result += int64(num)
		if num == 0 {
			result++
			zero2++
		}
	}

	if (zero1 == 0 && result > sum) || (zero2 == 0 && sum > result) {
		return -1
	}

	if sum > result {
		return sum
	}

	return result
}

func main() {
	// result: 12
	// nums1 := []int{3,2,0,1,0}
	// nums2 := []int{6,5,0}

	// result: -1
	nums1 := []int{2,0,2,0}
	nums2 := []int{1,4}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := minSum(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

