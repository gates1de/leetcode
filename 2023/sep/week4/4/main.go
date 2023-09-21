package main
import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	left := int(0)
	right := m

	for left <= right {
		partitionA := (left + right) / 2
		partitionB := (m + n + 1) / 2 - partitionA

		maxLeftA := math.MinInt32
		if partitionA > 0 {
			maxLeftA = nums1[partitionA - 1]
		}

		minRightA := math.MaxInt32
		if partitionA != m {
			minRightA = nums1[partitionA]
		}

		maxLeftB := math.MinInt32
		if partitionB > 0 {
			maxLeftB = nums2[partitionB - 1]
		}

		minRightB := math.MaxInt32
		if partitionB != n {
			minRightB = nums2[partitionB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (m + n) % 2 == 0 {
				return float64(max(maxLeftA, maxLeftB) + min(minRightA, minRightB)) / 2.0
			} else {
				return float64(max(maxLeftA, maxLeftB))
			}
		} else if maxLeftA > minRightB {
			right = partitionA - 1
		} else {
			left = partitionA + 1
		}
	}

	return float64(0)
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
	// result: 2.00000
	// nums1 := []int{1,3}
	// nums2 := []int{2}

	// result: 2.50000
	nums1 := []int{1,2}
	nums2 := []int{3,4}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

