package main
import (
	"fmt"
)

func singleNumber(nums []int) []int {
	m := make(map[int]bool)

	for _, n := range nums {
		if m[n] {
			delete(m, n)
		} else {
			m[n] = true
		}
	}

	result := make([]int, len(m))
	i := int(0)

	for n, _ := range m {
		result[i] = n
		i++
	}

	return result
}

func main() {
	// result: [3,5]
	// nums := []int{1, 2, 1, 3, 2, 5}

	// result: [-1,0]
	// nums := []int{-1, 0}

	// result: [1,0]
	// nums := []int{0, 1}

	// result: [1,3,7]
	nums := []int{3,9,8,4,2,7,6,4,2,6,9,8,1}

	// result: 
	// nums := []int{}

	result := singleNumber(nums)
	fmt.Printf("result = %v\n", result)
}

