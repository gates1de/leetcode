package main
import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	str := fmt.Sprintf("%v", strconv.FormatInt(int64(num), 2))
	result := int(0)
	for _, c := range str {
		if c == rune('1') {
			result++
		}
	}
	return result
}

func main() {
	// result: 3
	// num := uint32(11)

	// result: 1
	// num := uint32(128)

	// result: 31
	num := uint32(4294967293)

	// result: 
	// num := uint32(0)

	result := hammingWeight(num)
	fmt.Printf("result = %v\n", result)
}

