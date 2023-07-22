package main
import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	n := len(asteroids)
	stack := make([]int, 0, n)

	for _, asteroid := range asteroids {
		isNeededStack := true

		for len(stack) > 0 && stack[len(stack) - 1] > 0 && asteroid < 0 {
			peek := stack[len(stack) - 1]

			if abs(peek) < abs(asteroid) {
				stack = stack[:len(stack) - 1]
				continue
			} else if abs(peek) == abs(asteroid) {
				stack = stack[:len(stack) - 1]
			}

			isNeededStack = false
			break
		}

		if isNeededStack {
			stack = append(stack, asteroid)
		}
	}

	result := make([]int, len(stack))
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: [5,10]
	// asteroids := []int{5,10,-5}

	// result: []
	// asteroids := []int{8,-8}

	// result: [10]
	asteroids := []int{10,2,-5}

	// result: 
	// asteroids := []int{}

	result := asteroidCollision(asteroids)
	fmt.Printf("result = %v\n", result)
}

