package main
import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	digitsLen := len(digits)
	numOfDigit := int(0)
	for i := n; i > 0; i /= 10 {
		numOfDigit++
	}

	result := int(0)
	for i := 1; i < numOfDigit; i++ {
		cases := int(1)
		for j := 0; j < i; j++ {
			cases *= digitsLen
		}
		result += cases
	}

	helper("", digits, numOfDigit, n, &result)
	return result
}

func helper(current string, digits []string, maxDigit int, n int, result *int) {
	if len(current) == maxDigit {
		*result++
		return
	}

	nextNumOfDigit := len(current) + 1
	headNum := getNumFromHead(n, nextNumOfDigit)
	for i := 0; i < len(digits); i++ {
		nextCurrent := current + digits[i]
		nextCurrentNum, _ := strconv.Atoi(nextCurrent)
		// fmt.Printf("nextCurrentNum = %v, headNum = %v\n", nextCurrentNum, headNum)
		if nextCurrentNum > headNum {
			return
		}
		helper(nextCurrent, digits, maxDigit, n, result)
	}
}

func getNumFromHead(num int, digit int) int {
	numOfDigit := int(0)
	for i := num; i > 0; i /= 10 {
		numOfDigit++
	}

	if numOfDigit < digit || digit == 0 {
		return 0
	}

	for i := digit; i < numOfDigit; i++ {
		num /= 10
	}
	return num
}

func main() {
	// result: 20
	// digits := []string{"1", "3", "5", "7"}
	// n := int(100)

	// result: 29523
	// digits := []string{"1", "4", "9"}
	// n := int(1000000000)

	// result: 1
	// digits := []string{"7"}
	// n := int(8)

	// result: 0
	// digits := []string{"2"}
	// n := int(1)

	// result: 153391688
	digits := []string{"1","2","3","4","5","6","7","8"}
	n := int(940860624)

	// result: 
	// digits := []string{}
	// n := int(0)

	result := atMostNGivenDigitSet(digits, n)
	fmt.Printf("result = %v\n", result)
}

