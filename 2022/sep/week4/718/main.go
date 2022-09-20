package main
import (
	"fmt"
)

func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 < n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	result := int(0)
	for i := 0; i < n1 + n2 - 1; i++ {
		targetNums1 := nums1[max(0, i - n1 + 1):min(n1, i + 1)]
		targetNums2 := nums2[max(0, n2 - i - 1):min(n2, n1 + n2 - i - 1)]
		if len(targetNums1) > len(targetNums2) {
			targetNums1 = targetNums1[max(0, len(targetNums1) - len(targetNums2)):]
		}

		count := countMax(targetNums1, targetNums2)
		if result < count {
			result = count
		}
	}
	return result
}

// must be nums1 length == nums2 length
func countMax(nums1 []int, nums2 []int) int {
	result := int(0)
	count := int(0)
	for i, num1 := range nums1 {
		num2 := nums2[i]
		if num1 == num2 {
			count++
		} else {
			count = 0
		}

		if result < count {
			result = count
		}
	}
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// nums1 := []int{1,2,3,2,1}
	// nums2 := []int{3,2,1,4,7}

	// result: 5
	// nums1 := []int{0,0,0,0,0}
	// nums2 := []int{0,0,0,0,0}

	// result: 4
	// nums1 := []int{3,2,1,4,5,6,7,0,9,8}
	// nums2 := []int{4,5,6,7}

	// result: 0
	nums1 := []int{1}
	nums2 := []int{0}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := findLength(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

