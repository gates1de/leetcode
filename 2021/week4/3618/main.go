package main
import (
	"fmt"
)

// More fast & Use less than memory  (not my solution)
func concatenatedBinary(n int) int {
	ans := int(1)
	len := int(4)
	for i := 2; i <= n; i++ {
        if i == len {
			len *= 2
		}
		ans = (ans * len + i) % (1e9 + 7)
		fmt.Printf("current = %v\n", ans)
    }
    return ans
}

// Accepted
func mySolution2(n int) int {
	result := int(0)
	for i := 1; i <= n; i++ {
		digit := calcBinDigit(i)
		// fmt.Printf("i = %v, digit = %v\n", i, digit)
		result = (result * pow(2, digit) + i) % (1e9 + 7)
		// fmt.Printf("current = %v\n", result)
	}
	return result
}

// Timeout Limit Exceeded
func mySolution1(n int) int {
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
			num = powWithModulo(2, len(allBinString) - i - 1)
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

func calcBinDigit(n int) int {
	current := int(1)
	result := int(0)
	for current <= n {
		current *= 2
		result++
	}
	return result
}

func pow(n int, m int) int {
	result := int(1)
	for i := 0; i < m; i++ {
		result *= n
	}
	return result
}

func powWithModulo(n int, m int) int {
	result := int(1)
	for i := 0; i < m; i++ {
		result = n * (result % (1e9 + 7))
	}
	return result
}

func main() {
	// n := int(1)
	// n := int(2)
	// n := int(3)
	// n := int(4)
	// n := int(5)
	// n := int(6)
	// n := int(7)
	// n := int(8)
	// n := int(12)
	n := int(42)
	// n := int(8933)

	result := concatenatedBinary(n)
	// result := mySolution2(n)
	// result := mySolution1(n)
	fmt.Printf("result = %v\n", result)
}

