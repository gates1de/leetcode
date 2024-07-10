package main
import (
	"fmt"
)

func minOperations(logs []string) int {
	stack := make([]string, 0, len(logs))

	for _, log := range logs {
		if log == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
		} else if log != "./" {
			stack = append(stack, log)
		}
	}

	return len(stack)
}

func main() {
	// result: 2
	// logs := []string{"d1/","d2/","../","d21/","./"}

	// result: 3
	// logs := []string{"d1/","d2/","./","d3/","../","d31/"}

	// result: 0
	logs := []string{"d1/","../","../","../"}

	// result: 
	// logs := []string{}

	result := minOperations(logs)
	fmt.Printf("result = %v\n", result)
}

