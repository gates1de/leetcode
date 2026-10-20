package main

import (
	"fmt"
	"sort"
)

func sortableIntegers(nums []int) int {
	n := len(nums)
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	result := 0
	for k := 1; k <= n; k++ {
		if n%k != 0 {
			continue
		}

		lps := make([]int, k)
		sortable := true
		for start := 0; start < n; start += k {
			lps[0] = 0
			length := 0
			for i := 1; i < k; i++ {
				for length > 0 && nums[start+i] != nums[start+length] {
					length = lps[length-1]
				}
				if nums[start+i] == nums[start+length] {
					length++
				}
				lps[i] = length
			}

			matched := 0
			for i := 0; i < 2*k-1; i++ {
				value := sorted[start+i%k]
				for matched > 0 && value != nums[start+matched] {
					matched = lps[matched-1]
				}
				if value == nums[start+matched] {
					matched++
				}
				if matched == k {
					break
				}
			}
			if matched != k {
				sortable = false
				break
			}
		}
		if sortable {
			result += k
		}
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{3,1,2}

	// result: 0
	// nums := []int{7,6,5}

	// result: 3
	nums := []int{5, 8}

	// result: 0
	// nums := []int{}

	result := sortableIntegers(nums)
	fmt.Printf("result = %v\n", result)
}
