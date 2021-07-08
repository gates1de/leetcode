package main
import (
	"fmt"
	"reflect"
)

func findLength(nums1 []int, nums2 []int) int {
	if reflect.DeepEqual(nums1, nums2) {
		return len(nums1)
	}

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	windowLen := len(nums2)
	arrLen := len(nums1) + (windowLen - 1) * 2
	arr := make([]int, arrLen)
	for i, _ := range arr {
		if windowLen - 1 <= i && i < windowLen + len(nums1) - 1 {
			arr[i] = nums1[i - windowLen + 1]
		} else {
			arr[i] = -1
		}
	}
	// fmt.Printf("windowLen = %v, arr = %v\n", windowLen, arr)

	result := int(0)
	for i := 0; i <= arrLen - windowLen; i++ {
		target := arr[i:windowLen + i]
		count := windowMaxCount(target, nums2)
		if result < count {
			result = count
		}
	}
	return result
}

// Condition: len(nums1) == len(nums2)
func windowMaxCount(nums1 []int, nums2 []int) int {
	// fmt.Printf("nums1 = %v, nums2 = %v\n", nums1, nums2)
	n := len(nums1)
	count := int(0)
	result := int(0)
	for i := 0; i < n; i++ {
		num1 := nums1[i]
		num2 := nums2[i]
		if num1 < 0 || num2 < 0 {
			continue
		}
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

func main() {
	// result: 3
	// nums1 := []int{1, 2, 3, 2, 1}
	// nums2 := []int{3, 2, 1, 4, 7}

	// result: 5
	// nums1 := []int{0, 0, 0, 0, 0}
	// nums2 := []int{0, 0, 0, 0, 0}

	// result: 7
	// nums1 := []int{1, 2, 3, 2, 1, 4, 7, 8}
	// nums2 := []int{0, 2, 3, 2, 1, 4, 7, 8}

	// result: 5
	// nums1 := []int{1, 4, 7, 10, 8, 7, 6, 5, 4, 4, 4, 3, 2, 1}
	// nums2 := []int{8, 7, 6, 5, 4, 3, 2, 1, 4, 7, 10, 13, 14, 15, 16}

	// result: 0
	nums1 := []int{70, 39, 25, 40, 7}
	nums2 := []int{52, 20, 67, 5, 31}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := findLength(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

