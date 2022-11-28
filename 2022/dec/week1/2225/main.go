package main
import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	m := map[int]int{}

	for _, match := range matches {
		if _, ok := m[match[0]]; !ok {
			m[match[0]] = 0
		}
		m[match[1]]++
	}

	answer1 := make([]int, 0, len(matches))
	answer2 := make([]int, 0, len(matches))
	for player, loseCount := range m {
		if loseCount == 0 {
			answer1 = append(answer1, player)
		} else if loseCount == 1 {
			answer2 = append(answer2, player)
		}
	}

	sort.Ints(answer1)
	sort.Ints(answer2)
	return [][]int{answer1, answer2}
}

func main() {
	// result: [[1,2,10],[4,5,7,8]]
	// matches := [][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}}

	// result: [[1,2,5,6],[]]
	matches := [][]int{{2,3},{1,3},{5,4},{6,4}}

	// result: 
	// matches := [][]int{}

	result := findWinners(matches)
	fmt.Printf("result = %v\n", result)
}

