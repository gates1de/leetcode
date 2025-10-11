package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	sortedSpells := make([][2]int, n)
	for i := range sortedSpells {
		sortedSpells[i][0] = spells[i]
		sortedSpells[i][1] = i
	}

	sort.Slice(sortedSpells, func(a, b int) bool {
		return sortedSpells[a][0] < sortedSpells[b][0]
	})
	sort.Ints(potions)

	result := make([]int, n)
	potionIndex := m - 1
	for _, sortedSpell := range sortedSpells {
		spell := sortedSpell[0]
		index := sortedSpell[1]

		for potionIndex >= 0 && int64(spell * potions[potionIndex]) >= success {
			potionIndex -= 1
		}
		result[index] = m - (potionIndex + 1)
	}

	return result
}

func main() {
	// result: [4,0,3]
	// spells := []int{5,1,3}
	// potions := []int{1,2,3,4,5}
	// success := int64(7)

	// result: [2,0,2]
	spells := []int{3,1,2}
	potions := []int{8,5,8}
	success := int64(16)

	// result: []
	// spells := []int{}
	// potions := []int{}
	// success := int64(0)

	result := successfulPairs(spells, potions, success)
	fmt.Printf("result = %v\n", result)
}

