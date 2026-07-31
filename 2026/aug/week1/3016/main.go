package main

import (
	"fmt"
)

const (
	alphabetSize = 26
	keysPerPress = 8
)

func minimumPushes(word string) int {
	frequencies := [alphabetSize]int{}
	for _, character := range word {
		frequencies[character-'a']++
	}

	for index := range frequencies {
		maxIndex := index
		for candidate := range frequencies {
			if candidate > index && frequencies[candidate] > frequencies[maxIndex] {
				maxIndex = candidate
			}
		}
		frequencies[index], frequencies[maxIndex] = frequencies[maxIndex], frequencies[index]
	}

	result := int(0)
	for index, frequency := range frequencies {
		if frequency == 0 {
			break
		}

		result += (index/keysPerPress + 1) * frequency
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
