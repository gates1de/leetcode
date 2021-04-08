package main
import (
	"fmt"
)

func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A); i++ {
        if i - A[i] > 1 || i - A[i] < -1 {
			return false
		}
	}
    return true
}

// Slow & Use more memory
func mySolution(A []int) bool {
	globals := []int{}
	locals := []int{}
	for i, num1 := range A {
		if i + 1 < len(A) && num1 > A[i + 1] {
			locals = append(locals, num1)
		}

		for _, num2 := range A[i + 1:] {
			if num1 > num2 {
				globals = append(globals, num1)
			}
		}
	}
	fmt.Printf("globals = %v, locals = %v\n", globals, locals)
	return len(globals) == len(locals)
}

func main() {
	// A := []int{1, 0, 2}
	A := []int{1, 2, 0}
	// A := []int{}

	result := isIdealPermutation(A)
	fmt.Printf("result = %v\n", result)
}

