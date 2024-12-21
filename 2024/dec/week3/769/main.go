package main
import (
	"fmt"
)

func maxChunksToSorted(arr []int) int {
	result := int(0)
	maxNum := int(0)

	for i, num := range arr {
		maxNum = max(maxNum, num)
		if maxNum == i {
			result++
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// arr := []int{4,3,2,1,0}

	// result: 4
	arr := []int{1,0,2,3,4}

	// result: 
	// arr := []int{}

	result := maxChunksToSorted(arr)
	fmt.Printf("result = %v\n", result)
}

