package main
import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	for i, num := range numbers {
		m[num] = i
	}

	for i, num := range numbers {
		if index, ok := m[target - num]; ok {
			return []int{i + 1, index + 1}
		}
	}

	return []int{}
}

func main() {
	// result: [1,2]
	// numbers := []int{2, 7, 11, 15}
	// target := int(9)

	// result: [1,3]
	// numbers := []int{2, 3, 4}
	// target := int(6)

	// result: [1,2]
	numbers := []int{-1, 0}
	target := int(-1)

	// result: 
	// numbers := []int{}
	// target := int(0)

	result := twoSum(numbers, target)
	fmt.Printf("result = %v\n", result)
}

