package main
import (
	"fmt"
)

func numWaterBottles(numBottles int, numExchange int) int {
	return helper(numBottles, 0, numExchange, 0)
}

func helper(numBottles int, empty int, numExchange int, sum int) int {
	if numBottles == 0 {
		return sum
	}

	sum += numBottles
	newNumBottles := (numBottles + empty) / numExchange
	newEmpty := (numBottles + empty) % numExchange
	return helper(newNumBottles, newEmpty, numExchange, sum)
}

func main() {
	// result: 13
	// numBottles := int(9)
	// numExchange := int(3)

	// result: 19
	numBottles := int(15)
	numExchange := int(4)

	// result: 
	// numBottles := int(0)
	// numExchange := int(0)

	result := numWaterBottles(numBottles, numExchange)
	fmt.Printf("result = %v\n", result)
}

