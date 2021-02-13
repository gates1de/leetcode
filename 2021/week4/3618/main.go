package main
import (
	"fmt"
)

func concatenatedBinary(n int) int {
	allBinString := ""
	for i := 1; i <= n; i++ {
		binString := decimalToBinaryString(i)
		// fmt.Printf("binString = %v, all = %v\n", binString, allBinString)
		allBinString += binString
		// bin, _ := strconv.ParseInt(allBinString, 2, 0)
		// fmt.Printf("bin = %v\n", bin)
	}
	// fmt.Printf("allBinString = %v\n", allBinString)
	result := int(0)
	for i := len(allBinString) - 1; i >= 0; i-- {
		b := allBinString[i]
		num := int(0)
		if b == 49 {
			num = pow(2, len(allBinString) - i - 1)
		}
		result += num
		result = result % (1e9 + 7)
		// fmt.Printf("b = %v, num = %v, current result = %v\n", b, num, result)
	}
	return result
}

func decimalToBinaryString(n int) string {
	return fmt.Sprintf("%b", n)
}

func pow(n int, m int) int {
	result := int(1)
	for i := 0; i < m; i++ {
		result = n * (result % (1e9 + 7))
	}
	return result
}

func main() {
	// n := int(1)
	// n := int(3)
	// n := int(12)
	// n := int(42)
	n := int(8933)
	result := concatenatedBinary(n)
	fmt.Printf("result = %v\n", result)
}

