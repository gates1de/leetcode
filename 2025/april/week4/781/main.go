package main
import (
	"fmt"
)

func numRabbits(answers []int) int {
	m := make(map[int]int)

	for _, num := range answers {
		m[num]++
	}

	result := int(0)
	for answer, count := range m {
		groupSize := answer + 1
		numOfGroups := (count + groupSize - 1) / groupSize
		result += numOfGroups * groupSize
	}

	return result
}

func main() {
	// result: 5
	// answers := []int{1,1,2}

	// result: 11
	answers := []int{10,10,10}

	// result: 
	// answers := []int{}

	result := numRabbits(answers)
	fmt.Printf("result = %v\n", result)
}

