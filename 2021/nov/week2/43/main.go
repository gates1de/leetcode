package main
import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n1Digit := len(num1)
	n2Digit := len(num2)
	headN1, _ := strconv.Atoi(string(num1[0]))
	headN2, _ := strconv.Atoi(string(num2[0]))

	arrLen := int(0)
	headProduct := headN1 * headN2
	if headProduct >= 10 {
		arrLen += 2
	} else if headProduct == 0 {
		return "0"
	} else {
		arrLen += 1
	}
	arrLen += (n1Digit - 1) + (n2Digit - 1)
	// fmt.Printf("headN1 = %v, headN2 = %v, arrLen = %v\n", headN1, headN2, arrLen)

	arr := make([]int, arrLen)
	for n1Digit := 1; n1Digit <= len(num1); n1Digit++ {
		n1, _ := strconv.Atoi(string(num1[len(num1) - n1Digit]))
		for n2Digit := 1; n2Digit <= len(num2); n2Digit++ {
			n2, _ := strconv.Atoi(string(num2[len(num2) - n2Digit]))
			product := n1 * n2
			index := (n1Digit - 1) + (n2Digit - 1)
			// fmt.Printf("n1 = %v, n2 = %v, index = %v\n", n1, n2, index)
			if product >= 10 {
				arr[index + 1] += product / 10
				arr[index] += product % 10
			} else {
				arr[index] += product
			}

			if arr[index] >= 10 {
				arr[index] %= 10
				if index + 1 >= len(arr) {
					arr = append(arr, 1)
				} else {
					arr[index + 1]++
				}
			}
			if index + 1 < len(arr) && arr[index + 1] >= 10 {
				arr[index + 1] %= 10
				if index + 2 >= len(arr) {
					arr = append(arr, 1)
				} else {
					arr[index + 2]++
				}
			}
			// fmt.Printf("arr = %v\n", arr)
		}
	}
	// fmt.Printf("arr = %v\n", arr)
	result := ""
	for i := len(arr) - 1; i >= 0; i-- {
		result += strconv.Itoa(arr[i])
	}
	return result
}

func main() {
	// result: "6"
	// num1 := "2"
	// num2 := "3"

	// result: "56088"
	// num1 := "123"
	// num2 := "456"

	// result: "0"
	// num1 := "5"
	// num2 := "0"

	// result: "321710800778064"
	// num1 := "342876"
	// num2 := "938271564"

	// result: "121932631112635269"
	num1 := "123456789"
	num2 := "987654321"

	// result: ""
	// num1 := ""
	// num2 := ""

	result := multiply(num1, num2)
	fmt.Printf("result = %v\n", result)
}

