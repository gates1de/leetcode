package main
import (
	"fmt"
)

func maximumTripletValue(nums []int) int64 {
	result := int64(0)
	imax := int64(0)
	dmax := int64(0)

	for _, num := range nums {
		result = max(result, dmax * int64(num))
		dmax = max(dmax, imax - int64(num))
		imax = max(imax, int64(num))
	}

	return result
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 77
	// nums := []int{12,6,1,2,7}

	// result: 133
	// nums := []int{1,10,3,4,19}

	// result: 0
	nums := []int{1,2,3}

	// result: 
	// nums := []int{}

	result := maximumTripletValue(nums)
	fmt.Printf("result = %v\n", result)
}

