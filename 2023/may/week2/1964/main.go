package main
import (
	"fmt"
)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	lowestIncreasingSequenceLength := int(0)
	result := make([]int, n)
	lowestIncreasingSequences := make([]int, n)

	for i, height := range obstacles {
		index := binarySearch(height, lowestIncreasingSequenceLength, lowestIncreasingSequences)
		if index == lowestIncreasingSequenceLength {
			lowestIncreasingSequenceLength++
		}
		lowestIncreasingSequences[index] = height
		result[i] = index + 1
	}

	return result
}

func binarySearch(target int, right int, nums []int) int {
	if right == 0 {
		return 0
	}

	left := int(0)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	// result: [1,2,3,3]
	// obstacles := []int{1,2,3,2}

	// result: [1,2,1]
	// obstacles := []int{2,2,1}

	// result: [1,1,2,3,2,2]
	obstacles := []int{3,1,5,6,4,2}

	// result: 
	// obstacles := []int{}

	result := longestObstacleCourseAtEachPosition(obstacles)
	fmt.Printf("result = %v\n", result)
}

