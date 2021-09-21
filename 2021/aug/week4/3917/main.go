package main
import (
	"fmt"
	"strconv"
)

func complexNumberMultiply(num1 string, num2 string) string {
	nums1 := splitComplex(num1)
	nums2 := splitComplex(num2)
	// fmt.Printf("nums1 = %v, nums2 = %v\n", nums1, nums2)
	val1 := nums1[0] * nums2[0]
	val2 := nums1[0] * nums2[1]
	val3 := nums1[1] * nums2[0]
	val4 := nums1[1] * nums2[1] * -1
	num := val1 + val4
	compNum := val2 + val3
	return fmt.Sprintf("%v+%vi", num, compNum)
}

func splitComplex(s string) [2]int {
	isPlus := false
	numRunes := []rune{}
	compNumRunes := []rune{}
	for _, r := range s {
		if r == rune('+') {
			isPlus = true
			continue
		} else if r == rune('i') {
			break
		}

		if !isPlus {
			numRunes = append(numRunes, r)
		} else {
			compNumRunes = append(compNumRunes, r)
		}
	}
	numString := string(numRunes)
	compNumString := string(compNumRunes)
	num, _ := strconv.Atoi(numString)
	compNum, _ := strconv.Atoi(compNumString)
	return [2]int{num, compNum}
}

func main() {
	// result: "0+2i"
	// num1 := "1+1i"
	// num2 := "1+1i"

	// result: "0+-2i"
	// num1 := "1+-1i"
	// num2 := "1+-1i"

	// result: 3+-4i
	num1 := "-1+-2i"
	num2 := "1+2i"

	// result: 
	// num1 := ""
	// num2 := ""

	result := complexNumberMultiply(num1, num2)
	fmt.Printf("result = %v\n", result)
}

