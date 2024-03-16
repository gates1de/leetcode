package main
import (
	"fmt"
)

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	result := int(0)
	count := int(0)

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		if _, ok := m[count]; ok {
			result = max(result, i - m[count])
		} else {
			m[count] = i
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{0,1}

	// result: 2
	nums := []int{0,1,0}

	// result: 
	// nums := []int{}

	result := findMaxLength(nums)
	fmt.Printf("result = %v\n", result)
}

