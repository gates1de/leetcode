package main
import (
	"fmt"
	"sort"
)

func isPossible(target []int) bool {
	sort.Ints(target)
	targetSum := sum(target)
	for i := len(target) - 1; i >= 0; i-- {
		num := target[i]
		if num == 1 {
			continue
		}

		if targetSum - num == 0 || num - (targetSum - num) < 1 {
			return false
		}

		if targetSum - num == 1 {
			target[i] = 1
			targetSum = 2
		} else {
			target[i] = num % (targetSum - num)
			targetSum -= num - target[i]
		}

		sort.Ints(target)
		i++
	}

	return true
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: true
	// target := []int{9, 3, 5}

	// result: false
	// target := []int{1, 1, 1, 2}

	// result: true
	// target := []int{8, 5}

	// result: false
 	// target := []int{9, 9, 9}

	// result: true
	// target := []int{1, 1000000000}

	// result: true
	target := []int{5, 2}

	// result: 
	// target := []int{}

	result := isPossible(target)
	fmt.Printf("result = %v\n", result)
}

