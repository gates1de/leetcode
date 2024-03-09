package main
import (
	"fmt"
)

func getCommon(nums1 []int, nums2 []int) int {
	i := int(0)
	j := int(0)
	for i < len(nums1) && j < len(nums2) {
		n1 := nums1[i]
		n2 := nums2[j]

		if n1 == n2 {
			return n1
		}

		if n1 < n2 {
			i++
		}

		if n1 > n2 {
			j++
		}
	}

	return -1
}

func main() {
	// result: 2
	// nums1 := []int{1,2,3}
	// nums2 := []int{2,4}

	// result: 2
	nums1 := []int{1,2,3,6}
	nums2 := []int{2,3,4,5}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := getCommon(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

