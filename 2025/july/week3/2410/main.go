package main

import (
	"fmt"
	"sort"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	m := len(players)
	n := len(trainers)
	result := int(0)

	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && players[i] > trainers[j] {
			j++
		}

		if j < n {
			result++
			j++
		}
	}

	return result
}

func main() {
	// result: 2
	// players := []int{4,7,9}
	// trainers := []int{8,2,5,8}

	// result: 1
	players := []int{1,1,1}
	trainers := []int{10}

	// result: 
	// players := []int{}
	// trainers := []int{}

	result := matchPlayersAndTrainers(players, trainers)
	fmt.Printf("result = %v\n", result)
}

