package main
import (
	"fmt"
)

func makeEqual(words []string) bool {
	equalCount := len(words)
	if equalCount <= 1 {
		return true
	}

	counter := make(map[byte]int)
	for _, word := range words {
		for _, char := range word {
			counter[byte(char)]++
		}
	}

	for _, count := range counter {
		if count % equalCount != 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// words := []string{"abc","aabc","bc"}

	// result: false
	words := []string{"ab","a"}

	// result: 
	// words := []string{}

	result := makeEqual(words)
	fmt.Printf("result = %v\n", result)
}

