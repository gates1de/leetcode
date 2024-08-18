package main
import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	fiveBills := int(0)
	tenBills := int(0)

	for _, bill := range bills {
		if bill == 5 {
			fiveBills++
		} else if bill == 10 {
			if fiveBills > 0 {
				fiveBills--
				tenBills++
			} else {
				return false
			}
		} else {
			if tenBills > 0 && fiveBills > 0 {
				fiveBills--
				tenBills--
			} else if fiveBills >= 3 {
				fiveBills -= 3
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	// result: true
	// bills := []int{5,5,5,10,20}

	// result: false
	bills := []int{5,5,10,10,20}

	// result: 
	// bills := []int{}

	result := lemonadeChange(bills)
	fmt.Printf("result = %v\n", result)
}

