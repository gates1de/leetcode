package main
import (
	"fmt"
)

func findKthPositive(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return k
	}

	maxNum := arr[n - 1]
	arrIndex := int(0)
	for i := 1; i <= maxNum; i++ {
		if arr[arrIndex] == i {
			arrIndex++
		} else {
			k--
		}

		if k == 0 {
			return i
		}
	}

	for k > 0 {
		maxNum++
		k--
	}
	return maxNum
}

func main() {
	// result: 9
	// arr := []int{2,3,4,7,11}
	// k := int(5)

	// result: 6
	// arr := []int{1,2,3,4}
	// k := int(2)

	// result: 3
	arr := []int{5,6,8}
	k := int(3)

	// result: 
	// arr := []int{}
	// k := int(0)

	result := findKthPositive(arr, k)
	fmt.Printf("result = %v\n", result)
}

