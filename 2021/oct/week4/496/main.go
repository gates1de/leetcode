package main
import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	tmp := map[int]int{}
	m := map[int]int{}
	for _, num := range nums2 {
		for tmpNum, _ := range tmp {
			if tmpNum < num {
				m[tmpNum] = num
				delete(tmp, tmpNum)
			}
		}
		tmp[num] = 1
	}
	for i, num := range nums1 {
		if val, ok := m[num]; ok {
			result[i] = val
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	// result: [-1,3,-1]
	// nums1 := []int{4, 1, 2}
	// nums2 := []int{1, 3, 4, 2}

	// result: [3,-1]
	// nums1 := []int{2, 4}
	// nums2 := []int{1, 2, 3, 4}

	// result: [43,90,-1,61,43,43,61,-1]
	nums1 := []int{0, 14, 90, 36, 1, 6, 15, 61}
	nums2 := []int{4, 3, 2, 6, 1, 0, 43, 46, 83, 14, 90, 36, 15, 61}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := nextGreaterElement(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

