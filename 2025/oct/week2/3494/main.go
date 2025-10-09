package main

import (
	"fmt"
)

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	done := make([]int64, n + 1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if done[j + 1] < done[j] {
				done[j + 1] = done[j]
			}

			done[j + 1] += int64(skill[j]) * int64(mana[i])
		}

		for j := n - 1; j > 0; j-- {
			done[j] = done[j + 1] - int64(skill[j]) * int64(mana[i])
		}
	}

	return done[n]
}

func main() {
	// result: 110
	// skill := []int{1,5,2,4}
	// mana := []int{5,1,4,2}

	// result: 5
	// skill := []int{1,1,1}
	// mana := []int{1,1,1}

	// result: 21
	skill := []int{1,2,3,4}
	mana := []int{1,2}

	// result: 
	// skill := []int{}
	// mana := []int{}

	result := minTime(skill, mana)
	fmt.Printf("result = %v\n", result)
}

