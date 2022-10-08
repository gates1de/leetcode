package main
import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := nums[0] + nums[1] + nums[2]

    for i, num := range nums {
        left := i + 1
        right := len(nums) - 1

        for left < right {
            sum := num + nums[left] + nums[right]

            if target >= sum {
                left++
            } else {
                right--
            }

            if abs(target, sum) < abs(target, result){
                result = sum
            }
        }
    }

    return result
}

func abs(a, b int) int {
    if b > a {
        return b - a
    }
    return a - b
}

func main() {
	// result: 2
	// nums := []int{-1,2,1,-4}
	// target := int(1)

	// result: 0
	// nums := []int{0,0,0}
	// target := int(1)

	// result: 6
	// nums := []int{1,2,3,4,5}
	// target := int(4)

	// result: 4
	// nums := []int{1,2,8,9,-7,-6,-5,3,4,0}
	// target := int(4)

	// result: 3
	// nums := []int{1,1,1,0}
	// target := int(100)

	// result: 3
	nums := []int{1,1,1,1}
	target := int(4)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := threeSumClosest(nums, target)
	fmt.Printf("result = %v\n", result)
}

