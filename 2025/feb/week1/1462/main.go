package main
import (
	"fmt"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	isPrerequisites := make([][]bool, numCourses)
	for i, _ := range isPrerequisites {
		isPrerequisites[i] = make([]bool, numCourses)
	}

	for _, edge := range prerequisites {
		isPrerequisites[edge[0]][edge[1]] = true
	}

	for i := 0; i < numCourses; i++ {
		for j := 0; j < numCourses; j++ {
			for target := 0; target < numCourses; target++ {
				isPrerequisites[j][target] = isPrerequisites[j][target] ||
				(isPrerequisites[j][i] && isPrerequisites[i][target])
			}
		}
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		result[i] = isPrerequisites[query[0]][query[1]]
	}

	return result
}

func main() {
	// result: [false, true]
	// numCourses := int(2)
	// prerequisites := [][]int{{1,0}}
	// queries := [][]int{{0,1},{1,0}}

	// result: [false, false]
	// numCourses := int(2)
	// prerequisites := [][]int{}
	// queries := [][]int{{1,0},{0,1}}

	// result: [true, true]
	numCourses := int(3)
	prerequisites := [][]int{{1,2},{1,0},{2,0}}
	queries := [][]int{{1,0},{1,2}}

	// result: []
	// numCourses := int(0)
	// prerequisites := [][]int{}
	// queries := [][]int{}

	result := checkIfPrerequisite(numCourses, prerequisites, queries)
	fmt.Printf("result = %v\n", result)
}

