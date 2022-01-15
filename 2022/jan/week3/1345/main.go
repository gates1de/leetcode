package main
import (
	"fmt"
)

func minJumps(arr []int) int {
	if len(arr) <= 2 {
		return len(arr) - 1
	}

	m := map[int][]int{}
	for i, num := range arr {
		if m[num] == nil {
			m[num] = []int{}
		}
		m[num] = append(m[num], i)
	}

	result := len(arr) - 1
	helper(arr[0], 0, 0, m, arr, &result)
	return result
}

func helper(current int, index int, step int, m map[int][]int, arr []int, result *int) {
	// fmt.Printf("current = %v, index = %v, step = %v\n", current, index, step)
	if index == len(arr) - 1 {
		if step < *result {
			*result = step
		}
		return
	}

	if step >= *result {
		return
	}

	nums := m[current]

	for _, i := range nums {
		if index == i {
			continue
		}
		helper(arr[i], i, step + 1, m, arr, result)
	}

	if index > 0 {
		helper(arr[index - 1], index - 1, step + 1, m, arr, result)
	}
	if index < len(arr) - 1 {
		helper(arr[index + 1], index + 1, step + 1, m, arr, result)
	}
}

func main() {
	// result: 3
	// arr := []int{100,-23,-23,404,100,23,23,23,3,404}

	// result: 0
	// arr := []int{7}

	// result: 1
	// arr := []int{1,7}

	// result: 1
	// arr := []int{7,6,9,6,9,6,9,7}

	// result: 3
	// arr := []int{3,8,7,2,9,1,6,7,1,6,4,5,3,8,7,6,4,9,1,6,4,8,3,1,6,4,3,2,9,1,7,4,8,1}

	// result: 3
	arr := []int{11,22,7,7,7,7,7,7,7,22,13}

	// result: 
	// arr := []int{}

	result := minJumps(arr)
	fmt.Printf("result = %v\n", result)
}

