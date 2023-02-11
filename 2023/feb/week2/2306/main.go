package main
import (
	"fmt"
)

func distinctNames(ideas []string) int64 {
	groups := make(map[byte]map[string]bool, 26)
	for i := byte(0); i < 26; i++ {
		groups[i] = make(map[string]bool)
	}
	for _, idea := range ideas {
		groups[idea[0] - 97][string(idea[1:])] = true
	}
	result := int64(0)
	for i := byte(0); i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			numOfMutual := int(0)
			for idea, _ := range groups[i] {
				if groups[j][idea] {
					numOfMutual++
				}
			}
			result += int64(2 * (len(groups[i]) - numOfMutual) * (len(groups[j]) - numOfMutual))
		}
	}

	return result
}

func main() {
	// result: 6
	// ideas := []string{"coffee","donuts","time","toffee"}

	// result: 0
	ideas := []string{"lack","back"}

	// result: 
	// ideas := []string{}

	result := distinctNames(ideas)
	fmt.Printf("result = %v\n", result)
}

