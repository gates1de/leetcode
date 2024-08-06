package main
import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	freq := make([]int, 26)
	for _, char := range word {
		freq[char - 'a']++
	}

	sort.Ints(freq)
	sortedFreq := make([]int, 26)
	for i := 0; i < 26; i++ {
		sortedFreq[i] = freq[25 - i]
	}

	result := int(0)
	for i := 0; i < 26; i++ {
		if sortedFreq[i] == 0 {
			break
		}

		result += (i / 8 + 1) * sortedFreq[i]
	}

	return result
}

func main() {
	// result: 5
	// word := "abcde"

	// result: 12
	// word := "xyzxyzxyzxyz"

	// result: 24
	word := "aabbccddeeffgghhiiiiii"

	// result: 
	// word := ""

	result := minimumPushes(word)
	fmt.Printf("result = %v\n", result)
}

