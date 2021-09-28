package main
import (
	"fmt"
)

func sortArrayByParityII(nums []int) []int {
	length := len(nums)
	i := 0
	j := 1
	for i < length && j < length {
		if nums[i] % 2 == 0 {
			i += 2
		} else if nums[j] % 2 == 1 {
			j += 2
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

// Slow & User more memory
func mySolution(nums []int) []int {
	evens := []int{}
	odds := []int{}
	for _, num := range nums {
		if num % 2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	result := make([]int, len(nums))
	for i, _ := range nums {
		if i % 2 == 0 {
			result[i] = evens[len(evens) - 1]
			evens = evens[:len(evens) - 1]
		} else {
			result[i] = odds[len(odds) - 1]
			odds = odds[:len(odds) - 1]
		}
	}

	return result
}

func main() {
	// result: [4,5,2,7]
	// nums := []int{4, 2, 5, 7}

	// result: [2,3]
	// nums := []int{2, 3}

	// result: [4,1,2,5,8,3]
	nums := []int{4,2,1,8,5,3}

	// result: 
	// nums := []int{}

	result := sortArrayByParityII(nums)
	fmt.Printf("result = %v\n", result)
}

