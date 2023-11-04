package main
import (
	"fmt"
)

func getLastMoment(n int, left []int, right []int) int {
	result := int(0)

	for _, num := range left {
		result = max(result, num)
	}

	for _, num := range right {
		result = max(result, n - num)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// n := int(4)
	// left := []int{4,3}
	// right := []int{0,1}

	// result: 7
	// n := int(7)
	// left := []int{}
	// right := []int{0,1,2,3,4,5,6,7}

	// result: 7
	n := int(7)
	left := []int{0,1,2,3,4,5,6,7}
	right := []int{}

	// result: 
	// n := int(0)
	// left := []int{}
	// right := []int{}

	result := getLastMoment(n, left, right)
	fmt.Printf("result = %v\n", result)
}

