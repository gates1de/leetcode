package main
import (
	"fmt"
)

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	ptr1 := int(0)
	ptr2 := int(0)

	result := make([][]int, 0)
	for ptr1 < n1 && ptr2 < n2 {
		if nums1[ptr1][0] == nums2[ptr2][0] {
			result = append(result, []int{nums1[ptr1][0], nums1[ptr1][1] + nums2[ptr2][1]})
			ptr1++
			ptr2++
		} else if nums1[ptr1][0] < nums2[ptr2][0] {
			result = append(result, nums1[ptr1])
			ptr1++
		} else {
			result = append(result, nums2[ptr2])
			ptr2++
		}
	}

	for ptr1 < n1 {
		result = append(result, nums1[ptr1])
		ptr1++
	}

	for ptr2 < n2 {
		result = append(result, nums2[ptr2])
		ptr2++
	}

	return result
}

func main() {
	// result: [[1,6],[2,3],[3,2],[4,6]]
	// nums1 := [][]int{{1,2},{2,3},{4,5}}
	// nums2 := [][]int{{1,4},{3,2},{4,1}}

	// result: [[1,3],[2,4],[3,6],[4,3],[5,5]]
	nums1 := [][]int{{2,4},{3,6},{5,5}}
	nums2 := [][]int{{1,3},{4,3}}

	// result: []
	// nums1 := [][]int{}
	// nums2 := [][]int{}

	result := mergeArrays(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

