package main
import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" {
			if len(stack) <= 1 {
				break
			}
			num := stack[len(stack) - 2] + stack[len(stack) - 1]
			stack = append(stack[:len(stack) - 2], num)
		} else if token == "-" {
			if len(stack) <= 1 {
				break
			}
			num := stack[len(stack) - 2] - stack[len(stack) - 1]
			stack = append(stack[:len(stack) - 2], num)
		} else if token == "*" {
			if len(stack) <= 1 {
				break
			}
			num := stack[len(stack) - 2] * stack[len(stack) - 1]
			stack = append(stack[:len(stack) - 2], num)
		} else if token == "/" {
			if len(stack) <= 1 || stack[len(stack) - 1] == 0 {
				break
			}
			num := stack[len(stack) - 2] / stack[len(stack) - 1]
			stack = append(stack[:len(stack) - 2], num)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
		// fmt.Printf("stack = %v\n", stack)
	}
	if len(stack) == 0 {
		return 0
	}
	return stack[0]
}

func main() {
	// result: 9
	// tokens := []string{"2", "1", "+", "3", "*"}

	// result: 6
	// tokens := []string{"4", "13", "5", "/", "+"}

	// result: 22
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	// result: 
	// tokens := []string{}

	result := evalRPN(tokens)
	fmt.Printf("result = %v\n", result)
}

