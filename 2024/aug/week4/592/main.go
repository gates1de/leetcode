package main
import (
	"fmt"
	"regexp"
	"strconv"
)

func fractionAddition(expression string) string {
	num := int(0)
	denom := int(1)

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	nums := re.FindAllString(expression, -1)

	for i := 0; i < len(nums); i += 2 {
		currentNum, _ := strconv.Atoi(nums[i])
		currentDenom, _ := strconv.Atoi(nums[i + 1])

		num = num * currentDenom + currentNum * denom
		denom = denom * currentDenom
	}

	gcd := abs(findGCD(num, denom))

	num /= gcd
	denom /= gcd

	return fmt.Sprintf("%d/%d", num, denom)
}

func findGCD(a int, b int) int {
	if a == 0 {
		return b
	}
	return findGCD(b % a, a)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: "0/1"
	// expression := "-1/2+1/2"

	// result: "1/3"
	// expression := "-1/2+1/2+1/3"

	// result: "-1/6"
	expression := "1/3-1/2"

	// result: ""
	// expression := ""

	result := fractionAddition(expression)
	fmt.Printf("result = %v\n", result)
}

