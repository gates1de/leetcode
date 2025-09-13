package main

import (
	"fmt"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	cannotCommunications := make(map[int]bool)

	for _, friendship := range friendships {
		tmp := make(map[int]bool)
		canCommunicate := false

		for _, language := range languages[friendship[0] - 1] {
			tmp[language] = true
		}
		for _, language := range languages[friendship[1]-1] {
			if tmp[language] {
				canCommunicate = true
				break
			}
		}

		if !canCommunicate {
			cannotCommunications[friendship[0] - 1] = true
			cannotCommunications[friendship[1] - 1] = true
		}
	}

	maxCount := int(0)
	counts := make([]int, n + 1)

	for person := range cannotCommunications {
		for _, language := range languages[person] {
			counts[language]++
			maxCount = max(maxCount, counts[language])
		}
	}

	return len(cannotCommunications) - maxCount
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// n := int(2)
	// languages := [][]int{{1},{2},{1,2}}
	// friendships := [][]int{{1,2},{1,3},{2,3}}

	// result: 2
	n := int(3)
	languages := [][]int{{2},{1,3},{1,2},{3}}
	friendships := [][]int{{1,4},{1,2},{3,4},{2,3}}

	// result: 
	// n := int(0)
	// languages := [][]int{}
	// friendships := [][]int{}

	result := minimumTeachings(n, languages, friendships)
	fmt.Printf("result = %v\n", result)
}

