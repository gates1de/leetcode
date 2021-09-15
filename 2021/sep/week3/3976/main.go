package main
import (
	"fmt"
)

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	} else if len(arr) == 2 {
		if arr[0] == arr[1] {
			return 1
		}
		return 2
	}

	result := int(1)
	count := int(2)
	for i := 1; i < len(arr) - 1; i++ {
		// fmt.Printf("%v, %v, %v\n", arr[i - 1], arr[i], arr[i + 1])
		if arr[i - 1] == arr[i] && arr[i] == arr[i + 1] {
			count = 1
		}
		if (arr[i - 1] < arr[i] && arr[i] > arr[i + 1]) ||
			(arr[i - 1] > arr[i] && arr[i] < arr[i + 1]) {
			count++
		} else if arr[i - 1] == arr[i] && arr[i] == arr[i + 1] {
			count = 1
		} else {
			count = 2
		}

		if result < count {
			result = count
		}
	}
	return result
}

func main() {
	// result: 5
	// arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}

	// result: 2
	// arr := []int{4, 8, 12, 16}

	// result: 2
	// arr := []int{4, 8, 12}

	// result: 1
	// arr := []int{100}

	// result: 9
	// arr := []int{4,8,12,3,9,4,0,17,38,9,4,10,3,4,2,4,0,8,9,1,1,1,4,2,3,9,4,1}

	// result: 1
	arr := []int{1, 1, 1}

	// result: 
	// arr := []int{}

	result := maxTurbulenceSize(arr)
	fmt.Printf("result = %v\n", result)
}

