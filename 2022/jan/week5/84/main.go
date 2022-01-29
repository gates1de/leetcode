package main
import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	dp := map[int]int{}
	helper(0, heights[0], dp, heights)
	fmt.Println(dp)
	result := int(0)
	for _, rect := range dp {
		if result < rect {
			result = rect
		}
	}
	return result
}

func helper(steps int, minHeight int, dp map[int]int, heights []int) {
	n := len(heights)
	// fmt.Printf("steps = %v, minHeight = %v, n = %v\n", steps, minHeight, n)
	if n == 0 {
		if dp[0] < steps * minHeight {
			dp[0] = steps * minHeight
		}
		return
	}

	height := heights[0]
	if dp[n] <= height {
		dp[n] = height
	}

	steps++
	if minHeight > height {
		minHeight = height
	}

	rect := steps * minHeight

	if dp[n] <= rect {
		dp[n] = rect
	}

	// Continue step
	helper(steps, minHeight, dp, heights[1:])

	if height > 0 {
		// Start this step
		helper(1, height, dp, heights[1:])
	}
}

func main() {
	// result: 10
	// heights := []int{2, 1, 5, 6, 2, 3}

	// result: 4
	// heights := []int{2, 4}

	// result: 1
	// heights := []int{1}

	// result: 42
	// heights := []int{3,0,9,2,5,6,0,9,6,4,2,3,9,4,7,8,9,7,8,9,2,6,1}

	// result: 11
	// heights := []int{1,0,10,1,1,1,1,1,1,1,1,1,1,0}

	// result: 9
	// heights := []int{9,0}

	// result: 6
	// heights := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	// result: 2
	heights := []int{0, 2, 0}

	// result: 
	// heights := []int{}

	result := largestRectangleArea(heights)
	fmt.Printf("result = %v\n", result)
}

