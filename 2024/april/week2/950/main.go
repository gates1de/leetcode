package main
import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	result := make([]int, n)
	isSkip := false
	indexInDeck := int(0)
	indexInResult := int(0)

	sort.Ints(deck)

	for indexInDeck < n {
		if result[indexInResult] == 0 {
			if !isSkip {
				result[indexInResult] = deck[indexInDeck]
				indexInDeck++
			}

			isSkip = !isSkip
		}

		indexInResult = (indexInResult + 1) % n
	}

	return result
}

func main() {
	// result: [2,13,3,11,5,17,7]
	// deck := []int{17,13,11,2,3,5,7}

	// result: [1,1000]
	deck := []int{1,1000}

	// result: []
	// deck := []int{}

	result := deckRevealedIncreasing(deck)
	fmt.Printf("result = %v\n", result)
}

