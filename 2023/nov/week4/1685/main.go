package main
import (
	"fmt"
)

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	totalSum := int(0)
	for _, num := range nums {
		totalSum += num
	}

	leftSum := int(0)
	result := make([]int, n)
	for i, num := range nums {
		rightSum := totalSum - leftSum - num
		leftCount := i
		rightCount := n - 1 - i
		leftTotal := leftCount * num - leftSum
		rightTotal := rightSum - rightCount * num
		result[i] = leftTotal + rightTotal
		leftSum += num
	}

	return result
}

func main() {
	// result: [4,3,5]
	// nums := []int{2,3,5}

	// result: [24,15,13,15,21]
	nums := []int{1,4,6,8,10}

	// result: []
	// nums := []int{}

	result := getSumAbsoluteDifferences(nums)
	fmt.Printf("result = %v\n", result)
}

