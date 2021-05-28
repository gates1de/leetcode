package main
import (
	"fmt"
)

func maximumUniqueSubarray(nums []int) int {
    res := 0
    acc := 0
    set := map[int]bool{}
    l := -1
    for r := 0; r < len(nums); r++ {
        for set[nums[r]] {
            l++
            acc -= nums[l]
            delete(set, nums[l])
        }
        acc += nums[r]
        set[nums[r]] = true
        res = max(res, acc)
    }

    return res
}

func max(a, b int) int {
    if a < b { return b }
    return a
}

// Slow
func mySolution(nums []int) int {
	result := int(0)
	current := int(0)
	queue := []int{}
	for _, num := range nums {
		// fmt.Printf("current queue = %v, num = %v\n", queue, num)
		numIndex := findIndex(num, queue)
		if numIndex >= 0 {
			if numIndex + 1 < len(queue) {
				queue = queue[numIndex + 1:]
			} else {
				queue = []int{}
			}
			queue = append(queue, num)
			// fmt.Printf("next queue = %v\n", queue)

			if result < current {
				result = current
			}
			current = sum(queue)
			continue
		}
		queue = append(queue, num)
		current += num
		if result < current {
			result = current
		}
	}
	return result
}

func findIndex(target int, nums []int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 17
	// nums := []int{4, 2, 4, 5, 6}

	// result: 8
	// nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}

	// result: 28412
	nums := []int{79, 30, 4174, 9999, 30, 4176, 9998, 34, 1, 4174, 9999, 4176, 9998}

	// result: 
	// nums := []int{}

	result := maximumUniqueSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

