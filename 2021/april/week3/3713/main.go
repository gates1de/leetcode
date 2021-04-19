package main
import (
	"fmt"
)

func combinationSum4(nums []int, target int) int {
	list := make([]int, 0)
	for _, num := range nums {
		current := num
		helper(nums, target, &current, &list)
	}
	return len(list)
}

func helper(nums []int, target int, current *int, list *[]int) {
	// fmt.Printf("current = %v\n", *current)
	if *current >= target {
		if *current == target {
			*list = append(*list, *current)
		}
		return
	}

	for _, num := range nums {
		newCurrent := *current + num
		helper(nums, target, &newCurrent, list)
	}
}

func main() {
	// result: 7
	// nums := []int{1, 2, 3}
	// target := int(4)

	// result: 0
	// nums := []int{9}
	// target := int(3)

	// result: 
	nums := []int{4, 2, 1}
	target := int(32)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := combinationSum4(nums, target)
	fmt.Printf("result = %v\n", result)
}

