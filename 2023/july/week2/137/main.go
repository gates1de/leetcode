package main
import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	for num, count := range m {
		if count < 3 {
			return num
		}
	}

	return -1
}

func main() {
	// result: 3
	// nums := []int{2,2,3,2}

	// result: 99
	nums := []int{0,1,0,1,0,1,99}

	// result: 
	// nums := []int{}

	result := singleNumber(nums)
	fmt.Printf("result = %v\n", result)
}

