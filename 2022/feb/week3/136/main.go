package main
import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if m[num] {
			delete(m, num)
			continue
		}
		m[num] = true
	}

	for num, _ := range m {
		return num
	}

	return 0
}

func main() {
	// result: 1
	// nums := []int{2, 2, 1}

	// result: 4
	// nums := []int{4, 1, 2, 1, 2}

	// result: 1
	// nums := []int{1}

	// result: 3
	nums := []int{7,8,2,4,1,6,3,1,6,2,7,8,4}

	// result: 
	// nums := []int{}

	result := singleNumber(nums)
	fmt.Printf("result = %v\n", result)
}

