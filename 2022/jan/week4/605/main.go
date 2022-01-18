package main
import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 && n == 1 {
		return flowerbed[0] == 0
	}

	zeroCount := int(1)
	if flowerbed[1] == 1 {
		zeroCount = 0
	}
	for i, flower := range flowerbed {
		if flower == 1 {
			zeroCount = 0
			continue
		}

		zeroCount++

		if zeroCount == 3 {
			n--
			if i < len(flowerbed) -1 && flowerbed[i + 1] == 1 {
				zeroCount = 0
			} else {
				zeroCount = 1
			}
		}
	}

	if zeroCount == 2 {
		n--
	}

	return n <= 0
}

func main() {
	// result: true
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := int(1)

	// result: false
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := int(2)

	// result: false
	// flowerbed := []int{1, 0, 1, 0, 1, 0, 1, 0}
	// n := int(1)

	// result: true
	// flowerbed := []int{0}
	// n := int(1)

	// result: false
	// flowerbed := []int{1}
	// n := int(1)

	// result: true
	// flowerbed := []int{1, 0, 1, 0, 1, 0, 1, 0}
	// n := int(0)

	// result: false
	// flowerbed := []int{0, 0, 1, 0, 1}
	// n := int(2)

	// result: false
	// flowerbed := []int{0, 1, 0, 1, 0, 0}
	// n := int(2)

	// result: false
	// flowerbed := []int{1, 0, 0, 0, 0, 1}
	// n := int(2)

	// result: true
	// flowerbed := []int{1, 0, 0, 0, 0}
	// n := int(2)

	// result: true
	flowerbed := []int{0, 0, 0, 0, 1}
	n := int(2)

	// result: 
	// flowerbed := []int{}
	// n := int(0)

	result := canPlaceFlowers(flowerbed, n)
	fmt.Printf("result = %v\n", result)
}

