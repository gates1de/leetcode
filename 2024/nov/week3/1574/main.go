package main
import (
	"fmt"
)

func findLengthOfShortestSubarray(arr []int) int {
	right := len(arr) - 1
	for right > 0 && arr[right] >= arr[right - 1] {
		right--
	}

	result := right
	left := int(0)

	for left < right && (left == 0 || arr[left - 1] <= arr[left]) {
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}

		result = min(result, right - left - 1)
		left++
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// arr := []int{1,2,3,10,4,2,3,5}

	// result: 4
	// arr := []int{5,4,3,2,1}

	// result: 0
	arr := []int{1,2,3}

	// result: 
	// arr := []int{}

	result := findLengthOfShortestSubarray(arr)
	fmt.Printf("result = %v\n", result)
}

