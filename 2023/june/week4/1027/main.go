package main
import (
	"fmt"
)

func longestArithSeqLength(nums []int) int {
	result := int(0)
	dp := make(map[int]map[int]int)
	for right := 0; right < len(nums); right++ {
		dp[right] = make(map[int]int)
		for left := 0; left < right; left++ {
			diff := nums[left] - nums[right]
			val := dp[left][diff]
			if val == 0 {
				val = 1
			}
			dp[right][diff] = val + 1
			result = max(result, dp[right][diff])
		}
	}
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// nums := []int{3,6,9,12}

	// result: 3
	// nums := []int{9,4,7,2,10}

	// result: 4
	nums := []int{20,1,15,3,10,5,8}

	// result: 
	// nums := []int{}

	result := longestArithSeqLength(nums)
	fmt.Printf("result = %v\n", result)
}

