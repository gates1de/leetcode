package main
import (
	"fmt"
)

func countSeniors(details []string) int {
	result := int(0)

	for _, detail := range details {
		ageTens := detail[11] - '0'
		ageOnes := detail[12] - '0'
		age := ageTens * 10 + ageOnes

		if age > 60 {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// details := []string{"7868190130M7522","5303914400F9211","9273338290F4010"}

	// result: 0
	details := []string{"1313579440F2036","2921522980M5644"}

	// result: 
	// details := []string{}

	result := countSeniors(details)
	fmt.Printf("result = %v\n", result)
}

