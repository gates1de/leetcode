package main
import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	result := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if counts[i] == 0 {
			result[1] = i
		}
		if counts[i] == 2 {
			result[0] = i
		}
	}

	return result
}

func main() {
	// result: [2,3]
	// nums := []int{1,2,2,4}

	// result: [1,2]
	// nums := []int{1,1}

	// result: [2,1]
	// nums := []int{3,2,2}

	// result: [3,1]
	nums := []int{3,2,3,4,6,5}

	// result: 
	// nums := []int{}

	result := findErrorNums(nums)
	fmt.Printf("result = %v\n", result)
}

