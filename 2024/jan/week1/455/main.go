package main
import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result := int(0)
	i := int(0)
	j := int(0)
	for i < len(g) {
		if len(s) - 1 < j {
			return result
		}

		if s[j] < g[i] {
			j++
			continue
		}

		result++
		i++
		j++
	}

	return result
}

func main() {
	// result: 1
	// g := []int{1,2,3}
	// s := []int{1,1}

	// result: 2
	g := []int{1,2}
	s := []int{1,2,3}

	// result: 
	// g := []int{}
	// s := []int{}

	result := findContentChildren(g, s)
	fmt.Printf("result = %v\n", result)
}

