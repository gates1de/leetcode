package main
import (
	"fmt"
)

func validPartition(nums []int) bool {
	n := len(nums)
	dp := [3]bool{}
	dp[0] = true

	for i := 0; i < n; i++ {
		dpIndex := i + 1
		result := false
            if i > 0 && nums[i] == nums[i - 1] {
                result = result || dp[(dpIndex - 2) % 3]
            }
            if i > 1 && nums[i] == nums[i - 1] && nums[i] == nums[i - 2] {
                result = result || dp[(dpIndex - 3) % 3]
            }
            if i > 1 && nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2 {
                result = result || dp[(dpIndex - 3) % 3]
            }

            dp[dpIndex % 3] = result
	}

	return dp[n % 3]
}

func main() {
	// result: true
	// nums := []int{4,4,4,5,6}

	// result: false
	nums := []int{1,1,1,2}

	// result: 
	// nums := []int{}

	result := validPartition(nums)
	fmt.Printf("result = %v\n", result)
}

