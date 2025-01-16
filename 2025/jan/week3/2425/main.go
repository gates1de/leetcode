package main
import (
	"fmt"
)

func xorAllNums(nums1 []int, nums2 []int) int {
	xor1 := int(0)
	xor2 := int(0)
	n1 := len(nums1)
	n2 := len(nums2)

	if n2 % 2 != 0 {
		for _, num := range nums1 {
			xor1 ^= num
		}
	}

	if n1 % 2 != 0 {
		for _, num := range nums2 {
			xor2 ^= num
		}
	}

	return xor1 ^ xor2
}

func main() {
	// result: 13
	// nums1 := []int{2,1,3}
	// nums2 := []int{10,2,5,0}

	// result: 0
	nums1 := []int{1,2}
	nums2 := []int{3,4}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := xorAllNums(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

