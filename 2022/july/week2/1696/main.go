package main
import (
	"fmt"
	"math"
)

func maxResult(nums []int, k int) int {
	queue := []int{0}
    for i := 1; i < len(nums); i++ {
        if queue[0] + k < i {
            queue = queue[1:]
        }
        nums[i] += nums[queue[0]]
        for len(queue) > 0 && nums[queue[len(queue) - 1]] <= nums[i] {
            queue = queue[:len(queue) - 1]
        }
        queue = append(queue, i)
    }

    return nums[len(nums) - 1]
}

// Wrong Answer
func ngSolution(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	result := math.MinInt32
	helper(n - 1, 0, nums, k, &result)
	return result
}

func helper(index int, currentSum int, nums []int, k int, result *int) {
	fmt.Println(index, currentSum)
	if index < 0 {
		*result = max(*result, currentSum)
		return
	}

	currentSum += nums[index]
	if index == 0 {
		*result = max(*result, currentSum)
		return
	}

	headIndex := max(0, index - k)
	tailIndex := index

	targetNums := nums[headIndex:tailIndex]
	fmt.Println(targetNums)
	maxIndex := int(0)
	maxNumForMinus := int(-100001)
	for i := len(targetNums) - 1; i >= 0; i-- {
		num := targetNums[i]
		if num >= 0 {
			helper(index - (len(targetNums) - i), currentSum, nums, k, result)
			return
		}

		if maxNumForMinus <= num {
			maxIndex = i
			maxNumForMinus = num
		}
	}

	if headIndex == 0 {
		helper(0, currentSum, nums, k, result)
		return
	}
	helper(index - (len(targetNums) - maxIndex), currentSum, nums, k, result)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 7
	// nums := []int{1,-1,-2,4,-7,3}
	// k := int(2)

	// result: 17
 	// nums := []int{10,-5,-2,4,0,3}
	// k := int(3)

	// result: 0
 	// nums := []int{1,-5,-20,4,-1,3,-6,-3}
	// k := int(2)

	// result: -1
 	// nums := []int{0,-1,-2,-3,1}
	// k := int(2)

	// result: -40
 	nums := []int{40,30,-100,-100,-10,-7,-3,-3}
	k := int(2)

	// result: 
 	// nums := []int{}
	// k := int(0)

	result := maxResult(nums, k)
	fmt.Printf("result = %v\n", result)
}

