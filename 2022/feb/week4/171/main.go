package main
import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	letters := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
		'H': 8,
		'I': 9,
		'J': 10,
		'K': 11,
		'L': 12,
		'M': 13,
		'N': 14,
		'O': 15,
		'P': 16,
		'Q': 17,
		'R': 18,
		'S': 19,
		'T': 20,
		'U': 21,
		'V': 22,
		'W': 23,
		'X': 24,
		'Y': 25,
		'Z': 26,
	}

	result := int(0)
	digit := len(columnTitle)
	for _, l := range columnTitle {
		add := int(1)
		if digit > 1 {
			for i := 1; i < digit; i++ {
				add *= 26
			}
			add *= letters[l]
			result += add
		} else {
			result += letters[l]
		}
		digit--
	}

	return result
}

func main() {
	// result: 1
	// columnTitle := "A"

	// result: 28
	// columnTitle := "AB"

	// result: 701
	// columnTitle := "ZY"

	// result: 18279
	// columnTitle := "AAAA"

	// result: 214442
	// columnTitle := "LEET"

	// result: 2147483647
	columnTitle := "FXSHRXW"

	result := titleToNumber(columnTitle)
	fmt.Printf("result = %v\n", result)
}

