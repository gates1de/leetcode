package main
import (
	"fmt"
	"strconv"
	"strings"
)

func minMaxDifference(num int) int {
	numStr := strconv.Itoa(num)
	temp := numStr
	i := int(0)

	for i < len(numStr) && numStr[i] == '9' {
		i++
	}

	if i < len(numStr) {
		numStr = strings.ReplaceAll(numStr, string(numStr[i]), "9")
	}

	temp = strings.ReplaceAll(temp, string(temp[0]), "0")
	maxNum, _ := strconv.Atoi(numStr)
	minNum, _ := strconv.Atoi(temp)
	return maxNum - minNum
}

func main() {
	// result: 99009
	// num := int(11891)

	// result: 99
	num := int(90)

	// result: 
	// num := int(0)

	result := minMaxDifference(num)
	fmt.Printf("result = %v\n", result)
}

