package main
import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool, 1000)
	m2 := make(map[int]bool, 1000)
	result1 := make([]int, 0, 1000)
	result2 := make([]int, 0, 1000)

	for _, num := range nums1 {
		m1[num] = true
	}

	for _, num := range nums2 {
		if !m1[num] && !m2[num] {
			result2 = append(result2, num)
		}
		m2[num] = true
	}

	m1 = make(map[int]bool, 1000)
	for _, num := range nums1 {
		if !m1[num] && !m2[num] {
			result1 = append(result1, num)
		}
		m1[num] = true
	}

	return [][]int{result1, result2}
}

func main() {
	// result: [[1,3],[4,6]]
	// nums1 := []int{1,2,3}
	// nums2 := []int{2,4,6}

	// result: [[3],[]]
	nums1 := []int{1,2,3,3}
	nums2 := []int{1,1,2,2}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := findDifference(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

