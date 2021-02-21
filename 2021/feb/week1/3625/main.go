package main
import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	nums := fmt.Sprintf("%032b", num)
	result := int(0)
	for _, num := range nums {
		if num == 49 {
			result++
		}
	}
	return result
}

func main() {
	// num := "00000000000000000000000000001011"
	// num := "00000000000000000000000010000000"
	num := "11111111111111111111111111111101"

	u, err := strconv.ParseUint(num, 2, 32)
	if err != nil {
		panic(err)
	}
	result := hammingWeight(uint32(u))
	fmt.Printf("result = %v\n", result)
}

