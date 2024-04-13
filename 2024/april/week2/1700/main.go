package main
import (
	"fmt"
)

func countStudents(students []int, sandwiches []int) int {
	circleStudentCount := int(0)
	squareStudentCount := int(0)

	for _, student := range students {
		if student == 0 {
			circleStudentCount++
		} else {
			squareStudentCount++
		}
	}

	for _, sandwich := range sandwiches {
		if sandwich == 0 && circleStudentCount == 0 {
			return squareStudentCount
		}

		if sandwich == 1 && squareStudentCount == 0 {
			return circleStudentCount
		}

		if sandwich == 0 {
			circleStudentCount--
		} else {
			squareStudentCount--
		}
	}

	return 0
}

func main() {
	// result: 0
	// students := []int{1,1,0,0}
	// sandwiches := []int{0,1,0,1}

	// result: 3
	students := []int{1,1,1,0,0,1}
	sandwiches := []int{1,0,0,0,1,1}

	// result: 
	// students := []int{}
	// sandwiches := []int{}

	result := countStudents(students, sandwiches)
	fmt.Printf("result = %v\n", result)
}

