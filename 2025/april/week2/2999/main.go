package main
import (
	"fmt"
	"math"
	"strconv"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	startS := strconv.FormatInt(start - 1, 10)
	finishS := strconv.FormatInt(finish, 10)
	return helper(finishS, s, limit) - helper(startS, s, limit)
}

func helper(x string, s string, limit int) int64 {
	if len(x) < len(s) {
		return 0
	}

	if len(x) == len(s) {
		if x >= s {
			return 1
		}
		return 0
	}

	suffix := x[len(x) - len(s):]
	result := int64(0)
	preLen := len(x) - len(s)

	for i := 0; i < preLen; i++ {
		digit := int(x[i] - '0')

		if limit < digit {
			result += int64(math.Pow(float64(limit + 1), float64(preLen - i)))
			return result
		}

		result += int64(digit) * int64(math.Pow(float64(limit + 1), float64(preLen - 1- i)))
	}

	if suffix >= s {
		result++
	}

	return result
}

func main() {
	// result: 5
	// start := int64(1)
	// finish := int64(6000)
	// limit := int(4)
	// s := "124"

	// result: 2
	// start := int64(15)
	// finish := int64(215)
	// limit := int(6)
	// s := "10"

	// result: 0
	start := int64(1000)
	finish := int64(2000)
	limit := int(4)
	s := "3000"

	// result: 
	// start := int64(0)
	// finish := int64(0)
	// limit := int(0)
	// s := ""

	result := numberOfPowerfulInt(start, finish, limit, s)
	fmt.Printf("result = %v\n", result)
}

