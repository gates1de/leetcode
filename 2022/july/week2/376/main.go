package main
import (
	"fmt"
)

func wiggleMaxLength(nums []int) int {
	n := len(nums)
    if n <= 1 {
		return n
	}

    index := int(1)
    for index < n && nums[index] == nums[index - 1] {
		index++
	}

    if index == n {
		return 1
	}

    isBiggerThanPrev := nums[index] > nums[index - 1]
    result := int(1)
    prev := nums[index - 1]
    for i := index; i < n; i++ {
        if (isBiggerThanPrev && nums[i] > prev) || (!isBiggerThanPrev && nums[i] < prev) {
            isBiggerThanPrev = !isBiggerThanPrev
            result++
        }
        prev = nums[i]
    }
    return result
}

// Wrong Answer
func ngSolution(nums []int) int {
	uniqueNums := []int{nums[0]}
	pre := nums[0]
	for _, num := range nums[1:] {
		if pre == num {
			continue
		}
		pre = num
		uniqueNums = append(uniqueNums, num)
	}

	n := len(uniqueNums)
	fmt.Println(uniqueNums)
	if n <= 1 {
		return n
	}

	memo := make([]int, n + 1)

	helper(2, 1, uniqueNums[1] - uniqueNums[0] > 0, 2, uniqueNums, memo) 

	return memo[n]
}

func helper(index int, preIndex int, preIsUp bool, count int, nums []int, memo []int) {
	if index >= len(nums) {
		if count > memo[index] {
			memo[index] = count
		}
		return
	}

	if count < memo[index] {
		return
	}
	memo[index] = count

	preNum := nums[preIndex]
	num := nums[index]
	isUp := num - preNum > 0

	helper(index + 1, preIndex, preIsUp, count, nums, memo)
	if (preIsUp && !isUp) || (!preIsUp && isUp) {
		helper(index + 1, index, isUp, count + 1, nums, memo)
	}
}

func main() {
	// result: 6
	// nums := []int{1,7,4,9,2,5}

	// result: 7
	// nums := []int{1,17,5,10,13,15,10,5,16,8}

	// result: 2
	// nums := []int{1,2,3,4,5,6,7,8,9}

	// result: 4
	// nums := []int{1, 3, 4, 3, 4}

	// result: 1
	nums := []int{0, 0, 0}

	// result: 
	// nums := []int{}

	result := wiggleMaxLength(nums)
	fmt.Printf("result = %v\n", result)
}

