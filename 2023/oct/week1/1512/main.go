package main
import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	n := len(nums)
	sums := make(map[int]int)
	for i := 2; i <= n; i++ {
		sums[i] = sums[i - 1] + i - 1
	}

	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	result := int(0)
	for _, count := range counter {
		result += sums[count]
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{1,2,3,1,1,3}

	// result: 6
	// nums := []int{1,1,1,1}

	// result: 0
	nums := []int{1,2,3}

	// result: 
	// nums := []int{}

	result := numIdenticalPairs(nums)
	fmt.Printf("result = %v\n", result)
}

