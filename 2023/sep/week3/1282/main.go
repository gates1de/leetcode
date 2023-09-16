package main
import (
	"fmt"
)

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	result := make([][]int, 0, len(groupSizes))

	for i, num := range groupSizes {
		if m[num] == nil {
			m[num] = make([]int, 0, 3)
		}

		m[num] = append(m[num], i)
		if len(m[num]) == num {
			result = append(result, m[num])
			m[num] = nil
		}
	}

	for _, group := range m {
		if len(group) == 0 {
			continue
		}
		result = append(result, group)
	}
	return result
}

func main() {
	// result: [[5],[0,1,2],[3,4,6]]
	// groupSizes := []int{3,3,3,3,3,1,3}

	// result: [[1],[0,5],[2,3,4]]
	// groupSizes := []int{2,1,3,3,3,2}

	// result: [[0,2,3],[6,8,9],[1,4,5,7]]
	groupSizes := []int{3,4,3,3,4,4,3,4,3,3}

	// result: 
	// groupSizes := []int{}

	result := groupThePeople(groupSizes)
	fmt.Printf("result = %v\n", result)
}

