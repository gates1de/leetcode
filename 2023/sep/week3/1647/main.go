package main
import (
	"fmt"
)

func minDeletions(s string) int {
    counter := make(map[byte]int)
	result := int(0)
    usedFrequencies := make(map[int]bool)

    for _, c := range s {
        counter[byte(c)]++
    }

    for _, freq := range counter {
        for freq > 0 && usedFrequencies[freq] {
            freq--
			result++
        }
        usedFrequencies[freq] = true
    }

    return result
}

func main() {
	// result: 0
	// s := "aab"

	// result: 2
	// s := "aaabbbcc"

	// result: 2
	// s := "ceabaacb"

	// result: 1
	s := "accdcdadddbaadbc"

	// result: 
	// s := ""

	result := minDeletions(s)
	fmt.Printf("result = %v\n", result)
}

