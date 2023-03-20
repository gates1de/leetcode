package main
import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n <= 1 {
		return true
	}

	continuousZeroCount := int(0)

	for i, flower := range flowerbed {
		if flower == 0 {
			continuousZeroCount++
		} else {
			continuousZeroCount = 0
		}

		// can place head
		if i == 1 && continuousZeroCount == 2 {
			n--
			continuousZeroCount = 1
		}

		if continuousZeroCount == 3 {
			n--
			continuousZeroCount = 1
		}
	}

	// can place tail
	if continuousZeroCount == 2 {
		n--
	}

	return n <= 0
}

func main() {
	// result: true
	// flowerbed := []int{1,0,0,0,1}
	// n := int(1)

	// result: false
	// flowerbed := []int{1,0,0,0,1}
	// n := int(2)

	// result: true
	flowerbed := []int{0,0,1}
	n := int(1)

	// result: 
	// flowerbed := []int{}
	// n := int(0)

	result := canPlaceFlowers(flowerbed, n)
	fmt.Printf("result = %v\n", result)
}

