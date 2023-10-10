package main
import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	n := len(nums)
	result := n

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	newNums := make([]int, 0, len(m))
	for num, _ := range m {
		newNums = append(newNums, num)
	}

	sort.Ints(newNums)

	for i, num := range newNums {
		left := num
		right := left + n - 1
		j := binarySearch(newNums, right)
		count := j - i
		result = min(result, n - count)
	}

	return result
}

func binarySearch(newNums []int, target int) int {
	left := int(0)
	right := len(newNums)

	for left < right {
		mid := (left + right) / 2
		if target < newNums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// nums := []int{4,2,5,3}

	// result: 1
	// nums := []int{1,2,3,5,6}

	// result: 3
	nums := []int{1,10,100,1000}

	// result: 
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}

