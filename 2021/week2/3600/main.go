package main
import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		nums1 = nums2
		return
	} else if len(nums2) == 0 {
		nums1 = nums1[:m]
		return
	}

	copyNums1 := copySlice(nums1[:m])
	i := int(0)
	for len(copyNums1) > 0 || len(nums2) > 0 {
		var n1, n2 *int
		if len(copyNums1) > 0 {
			n1 = new(int)
			*n1 = copyNums1[0]
		}
		if len(nums2) > 0 {
			n2 = new(int)
			*n2 = nums2[0]
		}
		// fmt.Printf("n1 = %v, n2 = %v\n", *n1, *n2)

		if n1 == nil && n2 == nil {
			break
		} else if n1 != nil && n2 == nil {
			nums1[i] = *n1
			copyNums1 = copyNums1[1:]
		} else if n1 == nil && n2 != nil {
			nums1[i] = *n2
			nums2 = nums2[1:]
		} else if *n1 <= *n2 {
			nums1[i] = *n1
			copyNums1 = copyNums1[1:]
		} else if *n1 > *n2 {
			nums1[i] = *n2
			nums2 = nums2[1:]
		}
		i++
	}
}

func copySlice(target []int) []int {
	result := make([]int, len(target))
	copy(result, target)
	return result
}

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := int(3)
	// nums2 := []int{2,5,6}
	// n := int(3)

	// nums1 := []int{1}
	// m := int(1)
	// nums2 := []int{}
	// n := int(0)

	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := int(6)
	nums2 := []int{1, 2, 2}
	n := int(3)

	merge(nums1, m, nums2, n)
	fmt.Printf("result = %v\n", nums1)
}

