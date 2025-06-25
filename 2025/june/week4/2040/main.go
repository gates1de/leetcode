package main
import (
	"fmt"
	"math"
)

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n1 := len(nums1)
	n2 := len(nums2)
	pos1 := int(0)
	pos2 := int(0)

	for pos1 < n1 && nums1[pos1] < 0 {
		pos1++
	}
	for pos2 < n2 && nums2[pos2] < 0 {
		pos2++
	}

	left := int64(math.MinInt64)
	right := int64(math.MaxInt64)

	for left <= right {
		mid := (left + right) / 2
		count := int64(0)
		i := int(0)
		j := pos2 - 1

		for i < pos1 && j >= 0 {
			if int64(nums1[i]) * int64(nums2[j]) > mid {
				i++
			} else {
				count += int64(pos1 - i)
				j--
			}
		}

		i = pos1
		j = n2 - 1
		for i < n1 && j >= pos2 {
			if int64(nums1[i]) * int64(nums2[j]) > mid {
				j--
			} else {
				count += int64(j - pos2 + 1)
				i++
			}
		}

		i = 0
		j = pos2
		for i < pos1 && j < n2 {
			if int64(nums1[i]) * int64(nums2[j]) > mid {
				j++
			} else {
				count += int64(n2 - j)
				i++
			}
		}

		i = pos1
		j = 0
		for i < n1 && j < pos2 {
			if int64(nums1[i]) * int64(nums2[j]) > mid {
				i++
			} else {
				count += int64(n1 - i)
				j++
			}
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	// result: 8
	// nums1 := []int{2,5}
	// nums2 := []int{3,4}
	// k := int64(2)

	// result: 0
	// nums1 := []int{-4,-2,0,3}
	// nums2 := []int{2,4}
	// k := int64(6)

	// result: -6
	nums1 := []int{-2,-1,0,1,2}
	nums2 := []int{-3,-1,2,4,5}
	k := int64(3)

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}
	// k := int64(0)

	result := kthSmallestProduct(nums1, nums2, k)
	fmt.Printf("result = %v\n", result)
}

