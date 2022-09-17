package main
import (
	"fmt"
)

func validUtf8(data []int) bool {
	for i := 0; i < len(data); i++ {
		num := data[i]
		firstBits := decToBin(num)

		if firstBits[0] == byte('0') {
			continue
		}

		j := int(0)
		if firstBits[:5] == "11110" {
			// 4-byte
			j = 3
		} else if firstBits[:4] == "1110" {
			// 3-byte
			j = 2
		} else if firstBits[:3] == "110" {
			// 2-byte
			j = 1
		}
		
		if j == 0 {
			return false
		}

		if i + j > len(data) {
			return false
		}

		i++
		targetData := data[i:i + j]
		for _, b := range targetData {
			if decToBin(b)[:2] != "10" {
				return false
			}
		}

		i += j - 1
	}

	return true
}

func decToBin(n int) string {
	return fmt.Sprintf("%08b", n)
}

func main() {
	// result: true
	// data := []int{197,130,1}

	// result: false
	// data := []int{235,140,4}

	// result: true
	// data := []int{1}

	// result: false
	// data := []int{128}

	// result: true
	// data := []int{193,130,127,1}

	// result: false
	// data := []int{193,130,128,1}

	// result: true
	data := []int{228,189,160,229,165,189,13,10}

	// result: 
	// data := []int{}

	result := validUtf8(data)
	fmt.Printf("result = %v\n", result)
}

