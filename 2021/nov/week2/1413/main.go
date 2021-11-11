package main
import (
	"fmt"
)

func minStartValue(nums []int) int {
	sum := int(0)
	min := int(0)
	for _, num := range nums {
		sum += num
		if min > sum {
			min = sum
		}
	}
	return 1 - min
}

func main() {
	// result: 5
	// nums := []int{-3, 2, -3, 4, 2}

	// result: 1
	// nums := []int{1, 2}

	// result: 5
	// nums := []int{1, -2, -3}

	// result: 
	// nums := []int{}

	result := minStartValue(nums)
	fmt.Printf("result = %v\n", result)
}

