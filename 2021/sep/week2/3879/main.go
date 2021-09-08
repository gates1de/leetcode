package main
import (
	"fmt"
)

const byteA = byte(97)

func shiftingLetters(s string, shifts []int) string {
	for i, shift := range shifts {
		shifts[i] %= 26
		for j := 0; j < i; j++ {
			shifts[j] = (shifts[j] + shift) % 26
		}
	}
	// fmt.Printf("shifts = %v\n", shifts)

	chars := make([]rune, len(s))
	for i, r := range s {
		b := byte(r)
		shift := byte(shifts[i])
		newByte := byteA + ((b - byteA + shift) % 26)
		chars[i] = rune(newByte)
		// fmt.Printf("b = %v, newByte = %v, s = %v\n", b, newByte, string(rune(newByte)))
	}

	return string(chars)
}

func main() {
	// result: "rpl"
	// s := "abc"
	// shifts := []int{3, 5, 9}

	// result: "gfd"
	// s := "aaa"
	// shifts := []int{1, 2, 3}

	// result: "usoztmgjgkzuwmf"
	// s := "abcfdfjdwlbwalf"
	// shifts := []int{3,5,18,4,9,10,17,100,89,1,0,2,21,53,26}

	// result: "f"
	// s := "t"
	// shifts := []int{1000000000}

	// result: "nxokgpiejnscajwjnbtw"
	s := "benhwjsenjhvulyvefdn"
	shifts := []int{183265101,732053054,190062728,192602923,551817738,880105762,67914564,336769190,208588970,748586819,57544790,922070787,299351164,42155594,120233240,184503545,976640197,453293964,150570427,82070647}

	// result: 
	// s := ""
	// shifts := []int{}

	result := shiftingLetters(s, shifts)
	fmt.Printf("result = %v\n", result)
}

