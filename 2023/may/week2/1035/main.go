package main
import (
	"fmt"
)

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	dp := make([]int, n2 + 1)
	dpPrev := make([]int, n2 + 1)

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i - 1] == nums2[j - 1] {
				dp[j] = 1 + dpPrev[j - 1]
			} else {
				dp[j] = max(dp[j - 1], dpPrev[j])
			}
		}
		copy(dpPrev, dp)
	}

	return dp[n2]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums1 := []int{1,4,2}
	// nums2 := []int{1,2,4}

	// result: 3
	// nums1 := []int{2,5,1,2,5}
	// nums2 := []int{10,5,2,1,5,2}

	// result: 2
	nums1 := []int{1,3,7,1,7,5}
	nums2 := []int{1,9,2,5,1}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

