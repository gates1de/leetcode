package main
import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}
	if m == 0 {
		for i, num := range nums2 {
			nums1[i] = num
		}
		return
	}

	j := int(0)
	for i := 0; i < m; i++ {
		if nums1[i] < nums2[j] {
			continue
		}

		if nums1[i] >= nums2[j] {
			copy(nums1[i + 1:], nums1[i:m + n])
			nums1[i] = nums2[j]
		}

		j++
		if j >= n {
			break
		}
	}

	copy(nums1[m + j:], nums2[j:])
	sort.Ints(nums1)
}

func main() {
	// result: [1,2,2,3,5,6]
	// nums1 := []int{1,2,3,0,0,0}
	// m := int(3)
	// nums2 := []int{2,5,6}
	// n := int(3)

	// result: [1]
	// nums1 := []int{1}
	// m := int(1)
	// nums2 := []int{}
	// n := int(0)

	// result: [1]
	// nums1 := []int{0}
	// m := int(0)
	// nums2 := []int{1}
	// n := int(1)

	// result: [1,2]
	nums1 := []int{1,0}
	m := int(1)
	nums2 := []int{2}
	n := int(1)

	// result: 
	// nums1 := []int{}
	// m := int(0)
	// nums2 := []int{}
	// n := int(0)

	merge(nums1, m, nums2, n)
	fmt.Printf("result = %v\n", nums1)
}

