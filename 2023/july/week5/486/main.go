package main
import (
	"fmt"
)

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	copy(dp, nums)

	for diff := 1; diff < n; diff++ {
		for left := 0; left < n - diff; left++ {
			right := left + diff
			dp[left] = max(nums[left] - dp[left + 1], nums[right] - dp[left])
		}
	}

	if dp[0] < 0 {
		return false
	}

	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: false
	// nums := []int{1,5,2}

	// result: true
	nums := []int{1,5,233,7}

	// result: 
	// nums := []int{}

	result := PredictTheWinner(nums)
	fmt.Printf("result = %v\n", result)
}

