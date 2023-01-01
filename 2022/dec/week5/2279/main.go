package main
import (
	"fmt"
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
	remains := make([]int, n)
	for i, c := range capacity {
		rock := rocks[i]
		remains[i] = c - rock
	}
	sort.Ints(remains)

	for additionalRocks > 0 && len(remains) > 0 {
		fmt.Println(remains)
		remain := remains[0]
		if additionalRocks < remain {
			break
		} else if additionalRocks >= remain {
			remains = remains[1:]
			additionalRocks -= remain
		}
	}

	return n - len(remains)
}

func main() {
	// result: 3
	// capacity := []int{2,3,4,5}
	// rocks := []int{1,2,4,4}
	// additionalRocks := int(2)

	// result: 3
	capacity := []int{10,2,2}
	rocks := []int{2,2,0}
	additionalRocks := int(100)

	// result: 
	// capacity := []int{}
	// rocks := []int{}
	// additionalRocks := int(0)

	result := maximumBags(capacity, rocks, additionalRocks)
	fmt.Printf("result = %v\n", result)
}

