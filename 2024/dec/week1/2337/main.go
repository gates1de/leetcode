package main
import (
	"fmt"
)

func canChange(start string, target string) bool {
	n := len(start)
	startIndex := int(0)
	targetIndex := int(0)

	for startIndex < n || targetIndex < n {
			for startIndex < n && start[startIndex] == '_' {
				startIndex++
			}

			for targetIndex < n && target[targetIndex] == '_' {
				targetIndex++
			}

			if startIndex == n || targetIndex == n {
				return startIndex == n && targetIndex == n
			}

			startChar := start[startIndex]
			targetChar := target[targetIndex]
			if startChar != targetChar ||
				(startChar == 'L' && startIndex < targetIndex) ||
				(startChar == 'R' && startIndex > targetIndex) {
					return false
				}

				startIndex++
				targetIndex++
	}

	return true
}

func main() {
	// result: true
	// start := "_L__R__R_"
	// target := "L______RR"

	// result: false
	// start := "R_L_"
	// target := "__LR"

	// result: false
	start := "_R"
	target := "R_"

	// result: 
	// start := ""
	// target := ""

	result := canChange(start, target)
	fmt.Printf("result = %v\n", result)
}

