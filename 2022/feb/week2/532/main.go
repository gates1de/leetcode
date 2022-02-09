package main
import (
	"fmt"
)

func findPairs(nums []int, k int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	result := int(0)
	dup := map[[2]int]bool{}
	for num, count := range m {
		if count == 1 && (num + k == num || num - k == num) {
			continue
		}

		if m[num + k] > 0 && !dup[[2]int{num, num + k}] && !dup[[2]int{num + k, num}] {
			result++
			dup[[2]int{num, num + k}] = true
			dup[[2]int{num + k, num}] = true
		} else if m[num - k] > 0 && !dup[[2]int{num, num - k}] && !dup[[2]int{num - k, num}] {
			result++
			dup[[2]int{num, num - k}] = true
			dup[[2]int{num - k, num}] = true
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{3, 1, 4, 1, 5}
	// k := int(2)

	// result: 4
	// nums := []int{1, 2, 3, 4, 5}
	// k := int(1)

	// result: 1
	// nums := []int{1, 3, 1, 5, 4}
	// k := int(0)

	// result: 0
	// nums := []int{1}
	// k := int(1)

	// result: 8
	nums := []int{0,-3,2,9,-1,6,5,-8,9,-5,1,5,3,2,-6,0,1,4,3,2}
	k := int(5)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := findPairs(nums, k)
	fmt.Printf("result = %v\n", result)
}

