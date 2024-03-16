package main
import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = false
	}

	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			m[num] = true
		}
	}

	result := make([]int, 0, len(m))
	for num, isExists := range m {
		if isExists {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// result: [2]
	// nums1 := []int{1,2,2,1}
	// nums2 := []int{2,2}

	// result: [9,4]
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	// result: []
	// nums1 := []int{}
	// nums2 := []int{}

	result := intersection(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

