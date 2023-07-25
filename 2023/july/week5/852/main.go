package main
import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	l := int(0)
	r := len(arr) - 1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < arr[mid + 1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	// result: 1
	// arr := []int{0,1,0}

	// result: 1
	// arr := []int{0,2,1,0}

	// result: 1
	arr := []int{0,10,5,2}

	// result: 
	// arr := []int{}

	result := peakIndexInMountainArray(arr)
	fmt.Printf("result = %v\n", result)
}

