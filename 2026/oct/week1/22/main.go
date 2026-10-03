package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n < 0 {
		return []string{}
	}

	result := make([]string, 0)
	path := make([]byte, n*2)
	helper(path, 0, n, n, &result)
	return result
}

func helper(path []byte, position int, left int, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, string(path))
		return
	}

	if left > 0 {
		path[position] = '('
		helper(path, position+1, left-1, right, result)
	}
	if right > 0 && right > left {
		path[position] = ')'
		helper(path, position+1, left, right-1, result)
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
