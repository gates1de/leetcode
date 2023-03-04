package main
import (
	"fmt"
)

func compress(chars []byte) int {
	n := len(chars)
	if n <= 1 {
		return n
	}

	result := int(0)
	resultChars := make([]byte, 0, n)

	counter := int(1)
	pre := chars[0]
	for i := 1; i < n; i++ {
		char := chars[i]
		if pre == char {
			counter++
		} else {
			if counter == 1 {
				// alphabet
				result++
				resultChars = append(resultChars, pre)
			} else {
				// alphabet + digit count
				counterBytes := numToBytes(counter)
				result += 1 + len(counterBytes)
				resultChars = append(resultChars, pre)
				resultChars = append(resultChars, counterBytes...)
			}
			counter = 1
		}
		pre = char
	}

	if counter == 1 {
		result++
		resultChars = append(resultChars, pre)
	} else {
		counterBytes := numToBytes(counter)
		result += 1 + len(counterBytes)
		resultChars = append(resultChars, pre)
		resultChars = append(resultChars, counterBytes...)
	}
	copy(chars, resultChars)

	return result
}

func numToBytes(num int) []byte {
	result := make([]byte, 0, 32)
	for i := num; i > 0; i /= 10 {
		digit := i % 10
		digitChar := []byte(fmt.Sprintf("%v", digit))[0]
		result = append(result, '0')
		copy(result[1:], result[0:])
		result[0] = digitChar
	}
	return result
}

func main() {
	// result: 6
	// chars := []byte{'a','a','b','b','c','c','c'}

	// result: 1
	// chars := []byte{'a'}

	// result: 4
	chars := []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}

	// result: 
	// chars := []byte{}

	result := compress(chars)
	fmt.Printf("result = %v, chars = %v\n", result, chars[:result])
}

