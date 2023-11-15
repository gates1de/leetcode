package main
import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	result := int(1)
	for i := 1; i < len(arr); i++ {
		if arr[i] >= result + 1 {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// arr := []int{2,2,1,2,1}

	// result: 3
	// arr := []int{100,1,1000}

	// result: 5
	arr := []int{1,2,3,4,5}

	// result: 
	// arr := []int{}

	result := maximumElementAfterDecrementingAndRearranging(arr)
	fmt.Printf("result = %v\n", result)
}

