package main
import (
	"fmt"
	"math"
)

func minOperations(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	result := int(0)
	for _, count := range counter {
		if count == 1 {
			return -1
		}

		result += int(math.Ceil(float64(count) / 3))
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{2,3,3,2,2,4,2,3,4}

	// result: -1
	nums := []int{2,1,2,2,3,3}

	// result: 
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}

