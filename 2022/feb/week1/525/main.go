package main
import (
	"fmt"
)

func findMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	result := int(0)
	m := map[int]int{}
	countZero := int(0)
	countOne := int(0)
	for i, num := range nums {
		if num == 0 {
			countZero++
		} else {
			countOne++
		}

		if _, ok := m[countZero - countOne]; !ok {
			m[countZero - countOne] = i
		}

		if countZero == 0 || countOne == 0 {
			continue
		}

		subLen := int(0)
		if countZero == countOne {
			subLen = countZero * 2
		} else {
			if index, ok := m[countZero - countOne]; ok {
				subLen = i - index
			}
		}

		if result < subLen {
			result = subLen
		}
	}

    return result
}

func main() {
	// result: 2
	// nums := []int{0, 1}

	// result: 2
	// nums := []int{0, 1, 0}

	// result: 4
	// nums := []int{0, 1, 0, 1, 0}

	// result: 2
	// nums := []int{0, 1, 1, 1, 0}

	// result: 0
	// nums := []int{0}

	// result: 0
	// nums := []int{0, 0}

	// result: 8
	// nums := []int{1,1,1,1,1,1,0,0,0,0}

	// result: 32
	nums := []int{0,1,0,1,0,1,0,1,0,0,1,0,1,0,1,1,0,1,0,1,1,0,1,0,1,0,1,0,1,0,1,0,1,1,1,1,0,1,0,1,0,1,0}

	// result: 
	// nums := []int{}

	result := findMaxLength(nums)
	fmt.Printf("result = %v\n", result)
}

