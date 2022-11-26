package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	result := int(0)
	stack := make([]int, 0, n)

	for i := 0; i <= n; i++ {
		for len(stack) > 0 && (i == n || arr[stack[len(stack) - 1]] >= arr[i]) {
			mid := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			leftBoundary := int(-1)
			if len(stack) > 0 {
				leftBoundary = stack[len(stack) - 1]
			}
			rightBoundary := i

			count := ((mid - leftBoundary) * (rightBoundary - mid)) % modulo
			result += (count * arr[mid]) % modulo
			result %= modulo
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	// result: 17
	// arr := []int{3,1,2,4}

	// result: 444
	// arr := []int{11,81,94,43,3}

	// result: 525
	arr := []int{1,2,3,4,5,6,7,8,9,8,7,6,5,4,3,2,1}

	// result: 
	// arr := []int{}

	result := sumSubarrayMins(arr)
	fmt.Printf("result = %v\n", result)
}

