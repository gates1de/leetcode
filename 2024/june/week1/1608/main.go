package main
import (
	"fmt"
)

func specialArray(nums []int) int {
	n := len(nums)
	freq := make([]int, n + 1)

	for _, num := range nums {
		freq[min(n, num)]++
	}

	numGTOE := int(0)
	for i := n; i >= 1; i-- {
		numGTOE += freq[i]
		if i == numGTOE {
			return i
		}
	}

	return -1
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{3,5}

	// result: -1
	// nums := []int{0,0}

	// result: 3
	nums := []int{0,4,3,0,4}

	// result: 
	// nums := []int{}

	result := specialArray(nums)
	fmt.Printf("result = %v\n", result)
}

