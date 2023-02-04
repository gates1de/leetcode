package main
import (
	"fmt"
	"sort"
)

type Player struct {
	Score int
	Age int
}

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	players := make([]Player, n)
	for i, score := range scores {
		age := ages[i]
		players[i] = Player{Score: score, Age: age}
	}
	sort.Slice(players, func(a, b int) bool {
		if players[a].Age == players[b].Age {
			return players[a].Score < players[b].Score
		}
		return players[a].Age < players[b].Age
	})

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	return findMaxScore(players, -1, 0, dp)
}

func findMaxScore(players []Player, prev int, index int, dp [][]int) int {
	if index >= len(players) {
		return 0
	}

	if dp[prev + 1][index] > 0 {
		return dp[prev + 1][index]
	}

	if prev == -1 || players[index].Score >= players[prev].Score {
		dp[prev + 1][index] = max(
			findMaxScore(players, prev, index + 1, dp),
			players[index].Score + findMaxScore(players, index, index + 1, dp),
		)
	} else {
		dp[prev + 1][index] = findMaxScore(players, prev, index + 1, dp)
	}

	return dp[prev + 1][index]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 34
	// scores := []int{1,3,5,10,15}
	// ages := []int{1,2,3,4,5}

	// result: 16
	// scores := []int{4,5,6,5}
	// ages := []int{2,1,2,1}

	// result: 6
	scores := []int{1,2,3,5}
	ages := []int{8,9,10,1}

	// result: 
	// scores := []int{}
	// ages := []int{}

	result := bestTeamScore(scores, ages)
	fmt.Printf("result = %v\n", result)
}

