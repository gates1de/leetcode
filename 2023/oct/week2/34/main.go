package main
import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    } else if len(nums) == 1 {
        if nums[0] == target {
            return []int{0, 0}
        }
        return []int{-1, -1}
    }

    result := []int{-1, -1}
    head := int(0)
    tail := len(nums) - 1
    for head <= tail {
        mid := (head + tail) / 2
        midNum := nums[mid]

        if midNum == target {
            result[0] = mid
            result[1] = mid
            for i := mid; i >= 0; i-- {
                num := nums[i]
                result[0] = i
                if num != target {
                    result[0] = i + 1
                    break
                }
            }
            for i := mid; i < len(nums); i++ {
                num := nums[i]
                result[1] = i
                if num != target {
                    result[1] = i - 1
                    break
                }
            }
            break
        }
        if midNum < target {
            head = mid + 1
        } else {
            tail = mid - 1
        }
    }
    return result
}

func main() {
	// result: [3,4]
	// nums := []int{5,7,7,8,8,10}
	// target := int(8)

	// result: [-1,-1]
	// nums := []int{5,7,7,8,8,10}
	// target := int(6)

	// result: [-1,-1]
	nums := []int{}
	target := int(0)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := searchRange(nums, target)
	fmt.Printf("result = %v\n", result)
}

