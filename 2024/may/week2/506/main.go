package main
import (
	"fmt"
	"sort"
)

type Athlete struct {
	Order int
	Score int
}

func findRelativeRanks(score []int) []string {
	n := len(score)
	athletes := make([]Athlete, n)

	for i, s := range score {
		athletes[i] = Athlete{Order: i, Score: s}
	}

	sort.Slice(athletes, func(a, b int) bool {
		return athletes[a].Score > athletes[b].Score
	})

	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	result := make([]string, n)
	for i, athlete := range athletes {
		if i <= 2 {
			result[athlete.Order] = medals[i]
		} else {
			result[athlete.Order] = fmt.Sprintf("%d", i + 1)
		}
	}

	return result
}

func main() {
	// result: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
	// score := []int{5,4,3,2,1}

	// result: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
	score := []int{10,3,8,9,4}

	// result: []
	// score := []int{}

	result := findRelativeRanks(score)
	fmt.Printf("result = %v\n", result)
}

