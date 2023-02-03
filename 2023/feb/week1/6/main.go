package main
import (
	"fmt"
)

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}

	chars := make([][]byte, numRows)
	for i, _ := range chars {
		chars[i] = make([]byte, 0, n / 2)
	}

	row := int(0)
	charIndex := int(0)
	for charIndex < n {
		if charIndex % (numRows - 1) == 0 {
			for i := 0; i < numRows; i++ {
				if charIndex + i >= n {
					break
				}
				chars[i] = append(chars[i], s[charIndex + i])
			}
			charIndex += numRows 
			row = numRows - 2
		} else {
			chars[row] = append(chars[row], s[charIndex])
			charIndex++
			row--
		}
	}

	result := ""
	for _, bytes := range chars {
		result += string(bytes)
	}
	return result
}

func main() {
	// result: "PAHNAPLSIIGYIR"
	// s := "PAYPALISHIRING"
	// numRows := int(3)

	// result: "PINALSIGYAHRPI"
	// s := "PAYPALISHIRING"
	// numRows := int(4)

	// result: "A"
	s := "A"
	numRows := int(1)

	// result: 
	// s := ""
	// numRows := int(0)

	result := convert(s, numRows)
	fmt.Printf("result = %v\n", result)
}

