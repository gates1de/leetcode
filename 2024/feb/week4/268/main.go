package main
import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	sum := int(0)
	currentSum := int(0)
	isExistsZero := false
	for i, num := range nums {
		sum += i
		currentSum += num
		if n == num {
			currentSum -= num
		}
		if num == 0 {
			isExistsZero = true
		}
	}

	result := sum - currentSum
	if result == 0 {
		if !isExistsZero {
			return 0
		}
		return n
	}

	return sum - currentSum
}

func main() {
	// result: 2
	// nums := []int{3,0,1}

	// result: 2
	// nums := []int{0,1}

	// result: 8
	nums := []int{9,6,4,2,3,5,7,0,1}

	// result: 
	// nums := []int{}

	result := missingNumber(nums)
	fmt.Printf("result = %v\n", result)
}

