package main
import (
	"fmt"
	"math"
)

func maxDistance(s string, k int) int {
	result := int(0)
	north := int(0)
	south := int(0)
	east := int(0)
	west := int(0)

	for _, dir := range s {
			switch dir {
			case 'N':
					north++
			case 'S':
					south++
			case 'E':
					east++
			case 'W':
					west++
			}

			times1 := min(min(north, south), k)
			times2 := min(min(east, west), k - times1)
			current := count(north, south, times1) + count(east, west, times2)

			if current > result {
					result = current
			}
	}

	return result
}

func count(dir1 int, dir2, times int) int {
	return int(math.Abs(float64(dir1 - dir2))) + times * 2
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// s := "NWSE"
	// k := int(1)

	// result: 6
	s := "NSWWEW"
	k := int(3)

	// result: 
	// s := ""
	// k := int(0)

	result := maxDistance(s, k)
	fmt.Printf("result = %v\n", result)
}

