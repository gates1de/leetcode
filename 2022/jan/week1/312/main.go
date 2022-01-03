package main
import (
	"fmt"
)

func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

    fmt.Println(nums)
	n := len(nums)
    m := make([][]int, n)
    for i := range(m) {
        m[i] = make([]int, n)
    }

	for window := 2; window < n; window++ {
		for left := 0; left + window < n; left++ {
			max := int(0)
			right := left + window
			for i := left + 1; i < right; i++ {
				tmp := nums[left] * nums[i] * nums[right] + m[left][i] + m[i][right]
				if tmp > max {
					max = tmp
					m[left][right] = tmp
				}
			}
		}
	}

    return m[0][n - 1]
}

// Wrong Answer
func ngSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		result := nums[0] * nums[1]
		if nums[0] <= nums[1] {
			result += nums[1]
		} else {
			result += nums[0]
		}
		return result
	}

	result := int(0)
	newNums := []int{nums[0]}
	for i := 1; i < len(nums) - 1; i++ {
		prev := nums[i - 1]
		num := nums[i]
		next := nums[i + 1]
		if num <= 1 {
			if num == 1 {
				result += prev * next
			}
			continue
		}
		newNums = append(newNums, num)
	}
	newNums = append(newNums, nums[len(nums) - 1])

	// fmt.Printf("newNums = %v\n", newNums)
	helper(result, newNums, map[int]int{}, &result)

	return result
}

func helper(current int, nums []int, dp map[int]int, result *int) {
	// fmt.Printf("nums = %v\n", nums)
	if len(nums) == 0 {
		if *result < current {
			*result = current
		}
		return
	} else if len(nums) == 1 {
		current += nums[0]
		if *result < current {
			*result = current
		}
		return
	} else if len(nums) == 2 {
		current += nums[0] * nums[1]
		if nums[0] >= nums[1] {
			helper(current, []int{nums[0]}, dp, result)
		} else {
			helper(current, []int{nums[1]}, dp, result)
		}
		return
	}

	s := sum(nums)
	n := len(nums)
	if current < dp[s] {
		return
	}
	dp[s] = current

	prev := int(1)
	num := nums[0]
	next := nums[1]
	helper(current + (prev * num * next), nums[1:], dp, result)

	for i := 1; i < n - 1; i++ {
		prev := nums[i - 1]
		num := nums[i]
		next := nums[i + 1]

		newNums := make([]int, n - 1)
		copy(newNums, nums[:i])
		copy(newNums[i:], nums[i + 1:])
		helper(current + (prev * num * next), newNums, dp, result)
	}

	prev = nums[len(nums) - 2]
	num = nums[len(nums) - 1]
	next = int(1)
	helper(current + (prev * num * next), nums[:len(nums) - 1], dp, result)
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 167
	// nums := []int{3, 1, 5, 8}

	// result: 10
	// nums := []int{1, 5}

	// result: 724
	// nums := []int{3, 1, 5, 8, 9, 1, 2, 4}

	// result: 0
	// nums := []int{}

	// result: 2063 
	// nums := []int{3,2,8,1,6,9,3,2,7,8}

	// result: 279
	// nums := []int{2, 3, 7, 9, 1}

	// result: 1358
	// nums := []int{7, 9, 8, 0, 7, 1, 3, 5}

	// result: 3376
	nums := []int{1,6,8,3,4,6,4,7,9,8,0,6,2,8}

	// result: 
	// nums := []int{}

	result := maxCoins(nums)
	fmt.Printf("result = %v\n", result)
}

