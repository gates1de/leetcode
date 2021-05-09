package main
import (
	"fmt"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	result := 0
	l, r := sqrt(left, math.Ceil), sqrt(right, math.Floor)

	countSuperpalindromes := func(isEven bool) {
		for i := 1; ; i++ {
			p, j := i, i
			if !isEven {
				j /= 10
			}
			for ; j > 0; j /= 10 {
				p *= 10
				p += j % 10
			}
			if p > r {
				break
			}
			if p >= l && isPalindrome(p*p) {
				result++
			}
		}
	}

	countSuperpalindromes(true)
	countSuperpalindromes(false)

	return result
}

func isPalindrome(x int) bool {
	y := 0
	for ; x > y; x /= 10 {
		y *= 10
		y += x % 10
	}

	return x == y || x == y / 10
}

func sqrt(s string, roundFunc func(float64) float64) int {
	f, _ := strconv.ParseFloat(s, 64)
	return int(roundFunc(math.Sqrt(f)))
}

// Time Limit Exceeded
func ngSolution(left string, right string) int {
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
	if num < 10 {
		return true
	}
	copyNum := num
	reverse := int(0)
	currentDigit := int(1)
	// fmt.Printf("num = %v\n", num)
	for copyNum != 0 {
		reverse = reverse * 10 + copyNum % 10
		if reverse % 10 != copyNum % 10 {
			return false
		}
		// fmt.Printf("reverse = %v, %v, copyNum  = %v\n", reverse, reverse / currentDigit, copyNum % 10)
		copyNum /= 10
		currentDigit *= 10
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

