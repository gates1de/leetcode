package main
import (
	"fmt"
)

func winnerOfGame(colors string) bool {
	alice := int(0)
	bob := int(0)
	for i := 1; i < len(colors) - 1; i++ {
		if colors[i - 1] == colors[i] && colors[i] == colors[i + 1] {
			if colors[i] == 'A' {
				alice++
			} else {
				bob++
			}
		}
	}

	return alice - bob >= 1
}

func main() {
	// result: true
	// colors := "AAABABB"

	// result: false
	// colors := "AA"

	// result: false
	colors := "ABBBBBBBAAA"

	// result: 
	// colors := ""

	result := winnerOfGame(colors)
	fmt.Printf("result = %v\n", result)
}

