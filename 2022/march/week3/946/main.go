package main
import (
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	return helper([]int{pushed[0]}, pushed[1:], popped)
}

func helper(stack []int, pushed []int, popped []int) bool {
	// fmt.Println(stack, pushed, popped)
	if len(stack) == 0 {
		if len(pushed) == 0 {
			return true
		}
		stack = append(stack, pushed[0])
		pushed = pushed[1:]
		return helper(stack, pushed, popped)
	}

	stackNum := stack[len(stack) - 1]
	popNum := popped[0]
	if stackNum == popNum {
		if helper(stack[:len(stack) - 1], pushed, popped[1:]) {
			return true
		}
	}

	if len(pushed) == 0 {
		return false
	}
	stack = append(stack, pushed[0])
	pushed = pushed[1:]
	return helper(stack, pushed, popped)
}

func main() {
	// result: true
	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 5, 3, 2, 1}

	// result: false
	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 3, 5, 1, 2}

	// result: false
	// pushed := []int{1}
	// popped := []int{0}

	// result: true
	pushed := []int{10, 1, 11, 2, 12}
	popped := []int{10, 11, 12, 2, 1}

	// result: 
	// pushed := []int{}
	// popped := []int{}

	result := validateStackSequences(pushed, popped)
	fmt.Printf("result = %v\n", result)
}

