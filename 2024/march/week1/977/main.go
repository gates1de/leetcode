package main
import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, 0, n)
	i := int(0)
	j := n - 1
	for i <= j {
		head := abs(nums[i])
		tail := abs(nums[j])
		result = append(result, 0)

		if head > tail {
			copy(result[1:], result)
			result[0] = head * head
			i++
		} else {
			copy(result[1:], result)
			result[0] = tail * tail
			j--
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: [0,1,9,16,100]
	// nums := []int{-4,-1,0,3,10}

	// result: [4,9,9,49,121]
	nums := []int{-7,-3,2,3,11}

	// result: 
	// nums := []int{}

	result := sortedSquares(nums)
	fmt.Printf("result = %v\n", result)
}

