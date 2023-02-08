package main
import (
	"fmt"
)

func jump(nums []int) int {
    result := int(0)
    curEnd := int(0)
    curFarthest := int(0)

    for i := 0; i < len(nums) - 1; i++ {
        if curFarthest < i + nums[i] {
            curFarthest = i + nums[i]
        }
        if i == curEnd {
            result++
            curEnd = curFarthest
        }
    }
    return result
}

func main() {
	// result: 2
	// nums := []int{2,3,1,1,4}

	// result: 2
	nums := []int{2,3,0,1,4}

	// result: 
	// nums := []int{}

	result := jump(nums)
	fmt.Printf("result = %v\n", result)
}

