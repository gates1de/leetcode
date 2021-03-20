package main
import (
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) + len(popped) == 0 {
		return true
	} else if (len(pushed) + len(popped)) % 2 != 0 {
		return false
	}

	queue := []int{}
	for len(popped) > 0 {
		poppedNum := popped[0]
		// fmt.Printf("queue = %v, pushed = %v, popped = %v\n", queue, pushed, popped)
		if len(queue) == 0 {
			if len(pushed) == 0 {
				return false
			} else {
				queue = append(queue, pushed[0])
				pushed = pushed[1:]
			}
			continue
		}

		if queue[len(queue) - 1] != poppedNum {
			if len(pushed) > 0 {
				queue = append(queue, pushed[0])
				pushed = pushed[1:]
			} else {
				return false
			}
		} else {
			queue = queue[:len(queue) - 1]
			popped = popped[1:]
		}
	}
	return len(queue) == 0
}

func main() {
	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 5, 3, 2, 1}

	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 3, 5, 1, 2}

	// pushed := []int{}
	// popped := []int{}

	pushed := []int{1, 0}
	popped := []int{1, 0}

	result := validateStackSequences(pushed, popped)
	fmt.Printf("result = %v\n", result)
}

