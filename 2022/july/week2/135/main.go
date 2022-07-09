package main
import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}

	leftToRight := make([]int, n)
	rightToLeft := make([]int, n)

	leftToRight[0] = 1
	for i := 1; i < n; i++ {
		pre := ratings[i - 1]
		rating := ratings[i]
		if pre < rating {
			leftToRight[i] = leftToRight[i - 1] + 1
		} else {
			leftToRight[i] = 1
		}
	}

	rightToLeft[n - 1] = 1
	for i := n - 2; i >= 0; i-- {
		pre := ratings[i + 1]
		rating := ratings[i]
		if pre < rating {
			rightToLeft[i] = rightToLeft[i + 1] + 1
		} else {
			rightToLeft[i] = 1
		}
	}

	result := int(0)
	for i, left := range leftToRight {
		right := rightToLeft[i]
		result += max(left, right)
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
	// result: 5
	// ratings := []int{1, 0, 2}

	// result: 4
 	// ratings := []int{1, 2, 2}

	// result: 12
	// ratings := []int{29,51,87,87,72,12}

	// result: 1
	// ratings := []int{0}

	// result: 11
	ratings := []int{1,3,4,5,2}

	// result: 
	// ratings := []int{}

	result := candy(ratings)
	fmt.Printf("result = %v\n", result)
}

