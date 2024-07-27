package main
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	if n != len(heights) {
		return names
	}

	people := make([]Person, n)
	for i, name := range names {
		people[i] = Person{Name: name, Height: heights[i]}
	}

	sort.Slice(people, func (a, b int) bool {
		return people[a].Height > people[b].Height
	})

	result := make([]string, n)
	for i, person := range people {
		result[i] = person.Name
	}

	return result
}

func main() {
	// result: ["Mary","Emma","John"]
	// names := []string{"Mary","John","Emma"}
	// heights := []int{180,165,170}

	// result: ["Bob","Alice","Bob"]
	names := []string{"Alice","Bob","Bob"}
	heights := []int{155,185,150}

	// result: []
	// names := []string{}
	// heights := []int{}

	result := sortPeople(names, heights)
	fmt.Printf("result = %v\n", result)
}

