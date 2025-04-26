package main
import (
	"fmt"
)

func numberOfArrays(differences []int, lower int, upper int) int {
	x := int(0)
	y := int(0)
	current := int(0)

	for _, diff := range differences {
		current += diff
		x = min(x, current)
		y = max(y, current)

		if y - x > upper - lower {
			return 0
		}
	}

	return (upper - lower) - (y - x) + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// differences := []int{1,-3,4}
	// lower := int(1)
	// upper := int(6)

	// result: 4
	// differences := []int{3,-4,5,1,-2}
	// lower := int(-4)
	// upper := int(5)

	// result: 0
	differences := []int{4,-7,2}
	lower := int(3)
	upper := int(6)

	// result: 
	// differences := []int{}
	// lower := int(0)
	// upper := int(0)

	result := numberOfArrays(differences, lower, upper)
	fmt.Printf("result = %v\n", result)
}

