package main
import (
	"fmt"
)

func arraySign(nums []int) int {
	result := int(1)
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num > 0 {
			result *= 1
		} else {
			result *= -1
		}
	}
	return result
}

func main() {
	// result: 1
	// nums := []int{-1,-2,-3,-4,3,2,1}

	// result: 0
	// nums := []int{1,5,0,2,-3}

	// result: -1
	nums := []int{-1,1,-1,1,-1}

	// result: 
	// nums := []int{}

	result := arraySign(nums)
	fmt.Printf("result = %v\n", result)
}

