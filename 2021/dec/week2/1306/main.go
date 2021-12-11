package main
import (
	"fmt"
)

func canReach(arr []int, start int) bool {
	m := map[int]bool{}
	return helper(start, arr, m)
}

func helper(current int, arr []int, m map[int]bool) bool {
	// fmt.Printf("current = %v\n", current)
	if m[current] {
		return false
	}

	if arr[current] == 0 {
		return true
	}

	m[current] = true

	jump := arr[current]
	n := len(arr)
	isReached := false
	if current + jump < n {
		isReached = helper(current + jump, arr, m)
	}

	if isReached {
		return true
	}

	if current - jump >= 0 {
		isReached = helper(current - jump, arr, m)
	}

	return isReached
}

func main() {
	// result: true
	// arr := []int{4, 2, 3, 0, 3, 1, 2}
	// start := int(5)

	// result: true
	// arr := []int{4, 2, 3, 0, 3, 1, 2}
	// start := int(0)

	// result: false
	// arr := []int{3, 0, 2, 1, 2}
	// start := int(2)

	// result: true
	arr := []int{0, 1}
	start := int(1)

	// result: 
	// arr := []int{}
	// start := int(0)

	result := canReach(arr, start)
	fmt.Printf("result = %v\n", result)
}

