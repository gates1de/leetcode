package main
import (
	"fmt"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)
	result := int(0)
	i := int(1)
	for i * i <= rightNum {
		if i * i < leftNum {
			i++
			continue
		}
		if isPalindrome(i) && isPalindrome(i * i) {
			result++
		}
		i++
	}
	return result
}

func isPalindrome(num int) bool {
	copyNum := num
	reverse := int(0)
	for copyNum != 0 {
		reverse = reverse * 10 + copyNum % 10
		copyNum /= 10
	}

	return num == reverse
}

func main() {
	// result: 4
	// left := "4"
	// right := "1000"

	// result: 30
	left := "43143"
	right := "7072263972576"

	// result: 
	// left := ""
	// right := ""

	result := superpalindromesInRange(left, right)
	fmt.Printf("result = %v\n", result)
}

