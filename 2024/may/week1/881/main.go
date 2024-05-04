package main
import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	start := int(0)
	end := len(people) - 1
	result := int(0)

	for start <= end {
		if people[start] + people[end] <= limit {
			start++
		}

		end--
		result++
	}

	return result
}

func main() {
	// result: 1
	// people := []int{1,2}
	// limit := int(3)

	// result: 3
	// people := []int{3,2,2,1}
	// limit := int(3)

	// result: 4
	people := []int{3,5,3,4}
	limit := int(5)

	// result: 
	// people := []int{}
	// limit := int(0)

	result := numRescueBoats(people, limit)
	fmt.Printf("result = %v\n", result)
}

