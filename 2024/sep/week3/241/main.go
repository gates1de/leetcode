package main
import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	result := make([]int, 0)

	if len(expression) == 0 {
		return result
	}

	if len(expression) == 1 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
		return result
	}

	if len(expression) == 2 {
		if _, err := strconv.Atoi(string(expression[0])); err == nil {
      num, _ := strconv.Atoi(expression)
			result = append(result, num)
			return result
		}
	}

	for i := 0; i < len(expression); i++ {
		char := expression[i]

		if _, err := strconv.Atoi(string(char)); err == nil {
			continue
		}

		leftResults := diffWaysToCompute(expression[:i])
		rightResults := diffWaysToCompute(expression[i + 1:])

		for _, leftVal := range leftResults {
			for _, rightVal := range rightResults {
				computedResult := int(0)

				if char == '+' {
					computedResult = leftVal + rightVal
				} else if char == '-' {
					computedResult = leftVal - rightVal
				} else if char == '*' {
					computedResult = leftVal * rightVal
				}

				result = append(result, computedResult)
			}
		}
	}

	return result
}

func main() {
	// result: [0,2]
	// expression := "2-1-1"

	// result: [-34,-14,-10,-10,10]
	expression := "2*3-4*5"

	// result: []
	// expression := ""

	result := diffWaysToCompute(expression)
	fmt.Printf("result = %v\n", result)
}

