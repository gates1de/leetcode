package main
import (
	"fmt"
	"sort"
)

func divideArray(nums []int) bool {
	n := len(nums)
	if n % 2 != 0 || n == 0 {
		return false
	}

	sort.Ints(nums)
	pre := int(-1)
	for _, num := range nums {
		if pre == -1 {
			pre = num
		} else {
			if pre != num {
				return false
			}

			pre = -1
		}
	}

	return true
}

func main() {
	// result: true
	// nums := []int{3,2,3,2,2,2}

	// result: false
	nums := []int{1,2,3,4}

	// result: 
	// nums := []int{}

	result := divideArray(nums)
	fmt.Printf("result = %v\n", result)
}

