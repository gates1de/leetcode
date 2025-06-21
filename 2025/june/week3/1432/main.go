package main
import (
	"fmt"
	"strconv"
)

func maxDiff(num int) int {
	minNum := num
	maxNum := num

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			convertedStr := convert(num, x, y)

			if convertedStr[0] != '0' {
				temp, _ := strconv.Atoi(convertedStr)
				minNum = min(minNum, temp)
				maxNum = max(maxNum, temp)
			}
		}
	}

	return maxNum - minNum
}

func convert(num int, a int, b int) string {
	numStr := strconv.Itoa(num)
	result := make([]byte, len(numStr))

	for i, digit := range numStr {
		if int(digit - '0') == a {
			result[i] = byte('0' + b)
		} else {
			result[i] = byte(digit)
		}
	}

	return string(result)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 888
	// num := int(555)

	// result: 8
	num := int(9)

	// result: 
	// num := int(0)

	result := maxDiff(num)
	fmt.Printf("result = %v\n", result)
}

