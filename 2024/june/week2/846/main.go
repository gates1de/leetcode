package main
import (
	"fmt"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {
		return false
	}

	counter := make(map[int]int)
	for _, card := range hand {
		counter[card]++
	}

	for _, card := range hand {
		startCard := card

		for counter[startCard - 1] > 0 {
			startCard--
		}

		for startCard <= card {
			for counter[startCard] > 0 {
				for nextCard := startCard; nextCard < startCard + groupSize; nextCard++ {
					if counter[nextCard] == 0 {
						return false
					}

					counter[nextCard]--
				}
			}

			startCard++
		}
	}

	return true
}

func main() {
	// result: true
	// hand := []int{1,2,3,6,2,3,4,7,8}
	// groupSize := int(3)

	// result: false
	hand := []int{1,2,3,4,5}
	groupSize := int(4)

	// result: 
	// hand := []int{}
	// groupSize := int(0)

	result := isNStraightHand(hand, groupSize)
	fmt.Printf("result = %v\n", result)
}

