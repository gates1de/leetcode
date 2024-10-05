package main
import (
	"fmt"
)

func dividePlayers(skill []int) int64 {
	n := len(skill)
	totalSkill := int(0)
	m := make(map[int]int)

	for _, s := range skill {
		totalSkill += s
		m[s]++
	}

	if totalSkill % (n / 2) != 0 {
			return -1
	}

	targetSkill := totalSkill / (n / 2)
	result := int64(0)

	for currentSkill, val := range m {
		currentFreq := val
		partnerSkill := targetSkill - currentSkill

		if _, exists := m[partnerSkill]; !exists {
			return -1
		}
		if currentFreq != m[partnerSkill] {
			return -1
		}

		result += int64(currentSkill) * int64(partnerSkill) * int64(currentFreq)
	}

	return result / 2
}

func main() {
	// result:  22
	// skill := []int{3,2,5,1,3,4}

	// result: 12
	// skill := []int{3,4}

	// result: -1
	skill := []int{1,1,2,3}

	// result: 
	// skill := []int{}

	result := dividePlayers(skill)
	fmt.Printf("result = %v\n", result)
}

