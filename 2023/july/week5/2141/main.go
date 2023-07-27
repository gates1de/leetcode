package main
import (
	"fmt"
)

func maxRunTime(n int, batteries []int) int64 {
	sumPower := int64(0)
	for _, power := range batteries {
		sumPower += int64(power)
	}

	left := int64(0)
	right := sumPower / int64(n)
	for left < right {
		target := right - (right - left) / 2
		extra := int64(0)

		for _, power := range batteries {
			extra += min(int64(power), target)
		}

		if extra >= int64(n) * target {
			left = target
		} else {
			right = target - 1
		}
	}

	return left
}

func min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// n := int(2)
	// batteries := []int{3,3,3}

	// result: 2
	n := int(2)
	batteries := []int{1,1,1,1}

	// result: 
	// n := int(0)
	// batteries := []int{}

	result := maxRunTime(n, batteries)
	fmt.Printf("result = %v\n", result)
}

