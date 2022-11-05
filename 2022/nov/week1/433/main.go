package main
import (
	"fmt"
)

func minMutation(start string, end string, bank []string) int {
	levels := map[string]int{}
	queue := []string{start}
	for len(queue) > 0 {
		currentGene := queue[0]
		currentLevel := levels[currentGene]
		queue = queue[1:]

		for _, gene := range bank {
			level := levels[gene]
			isValid := valid(currentGene, gene)

			if isValid && gene == end {
				return currentLevel + 1
			}

			if isValid && (level == 0 || currentLevel + 1 <= level) {
				queue = append(queue, gene)
				levels[gene] = currentLevel + 1
			}
		}
	}

	return -1
}

func valid(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aChars := []byte(a)
	bChars := []byte(b)

	isExistsUnmatch := false
	for i, aChar := range aChars {
		bChar := bChars[i]
		if aChar != bChar {
			if isExistsUnmatch {
				return false
			}
			isExistsUnmatch = true
		}
	}
	return true
}

func main() {
	// result: 1
	// start := "AACCGGTT"
	// end := "AACCGGTA"
	// bank := []string{"AACCGGTA"}

	// result: 2
	// start := "AACCGGTT"
	// end := "AAACGGTA"
	// bank := []string{"AACCGGTA","AACCGCTA","AAACGGTA"}

	// result: 3
	// start := "AAAAACCC"
	// end := "AACCCCCC"
	// bank := []string{"AAAACCCC","AAACCCCC","AACCCCCC"}

	// result: 5
	start := "AAAAACCC"
	end := "AACCCCTT"
	bank := []string{"AAAACCCC","AAACCCCC","AACCCCCC","AAAACCCC","AACCCCCT","AACCCCTT"}

	// result: 
	// start := ""
	// end := ""
	// bank := []string{}

	result := minMutation(start, end, bank)
	fmt.Printf("result = %v\n", result)
}

