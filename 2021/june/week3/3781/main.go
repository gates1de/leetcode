package main
import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := []string{}
	helper("", n, n, &result)
	return result
}

func helper(current string, left int, right int, result *[]string) {
	// fmt.Printf("current = %v, left = %v, right = %v\n", current, left, right)
	if left == 0 && right == 0 {
		*result = append(*result, current)
		// fmt.Printf("result = %v\n", *result)
		return
	}

	if left > 0 {
		helper(current + "(", left - 1, right, result)
	}
	if right > 0 && right > left {
		helper(current + ")", left, right - 1, result)
	}
}

func main() {
	// result: ["((()))","(()())","(())()","()(())","()()()"]
	// n := int(3)

	// result: ["()"]
	n := int(1)

	// result: 
	// n := int()

	result := generateParenthesis(n)
	fmt.Printf("result = %v\n", result)
}

