package main
import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	lastAdded := int(1000)
	index := int(0)
	for _, num := range nums {
		if lastAdded == num {
			continue
		}

		nums[index] = num
		index++
		lastAdded = num
	}

	return index
}

func main() {
	// result: 2, [1,2,_]
	// nums := []int{1,1,2}

	// result: 5, [0,1,2,3,4,_,_,_,_,_]
	nums := []int{0,0,1,1,1,2,2,3,3,4}

	// result: 
	// nums := []int{}

	result := removeDuplicates(nums)
	fmt.Printf("result = %v\n", result)
}

