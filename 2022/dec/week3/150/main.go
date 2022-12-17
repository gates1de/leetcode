package main
import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 10000)

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			if len(stack) < 2 {
				// tokens is not valid
				continue
			}

			a := stack[len(stack) - 2]
			b := stack[len(stack) - 1]
			val := calc(a, b, []byte(token)[0])
			stack = stack[:len(stack) - 1]
			stack[len(stack) - 1] = val
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	if len(stack) == 0 {
		return 0
	}

	return stack[0]
}

func calc(a int, b int, operand byte) int {
	switch operand {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			return 0
		}
		return a / b
	}
	return 0
}

func main() {
	// result: 9
	// tokens := []string{"2","1","+","3","*"}

	// result: 6
	// tokens := []string{"4","13","5","/","+"}

	// result: 22
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	// result: 
	// tokens := []string{}

	result := evalRPN(tokens)
	fmt.Printf("result = %v\n", result)
}

