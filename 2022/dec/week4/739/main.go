package main
import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < temperatures[i] {
			result[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return result
}

func main() {
	// result: [1,1,4,2,1,1,0,0]
	// temperatures := []int{73,74,75,71,69,72,76,73}

	// result: [1,1,1,0]
	// temperatures := []int{30,40,50,60}

	// result: [1,1,0]
	temperatures := []int{30,60,90}

	// result: 
	// temperatures := []int{}

	result := dailyTemperatures(temperatures)
	fmt.Printf("result = %v\n", result)
}

