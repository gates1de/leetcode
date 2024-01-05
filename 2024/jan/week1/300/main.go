package main
import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    dp := make([]int, n)
    dp[0] = nums[0]
    tail := int(0)

    for i := 1; i < len(nums); i++ {
        index := binarySearchIndex(nums[i], 0, tail, dp)
        dp[index] = nums[i]
        if index == tail + 1 {
            tail++
        }
    }

    return tail + 1
}

func binarySearchIndex(target int, head int, tail int, dp []int) int {
    for head <= tail {
        mid := head + (tail - head) / 2
        if dp[mid] > target {
            tail = mid - 1
        } else if dp[mid] < target {
            head = mid + 1
        } else {
            return mid
        }
    }

    return head
}

func main() {
	// result: 4
	// nums := []int{10,9,2,5,3,7,101,18}

	// result: 4
	// nums := []int{0,1,0,3,2,3}

	// result: 1
	nums := []int{7,7,7,7,7,7,7}

	// result: 
	// nums := []int{}

	result := lengthOfLIS(nums)
	fmt.Printf("result = %v\n", result)
}

