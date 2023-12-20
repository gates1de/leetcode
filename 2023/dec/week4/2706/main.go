package main
import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {
	n := len(prices)
	if n <= 1 {
		return money
	}

	sort.Ints(prices)

	choco1 := prices[0]
	choco2 := prices[1]
	if choco1 + choco2 > money {
		return money
	}

	return money - choco1 - choco2
}

func main() {
	// result: 0
	// prices := []int{1,2,2}
	// money := int(3)

	// result: 3
	prices := []int{3,2,3}
	money := int(3)

	// result: 
	// prices := []int{}
	// money := int(0)

	result := buyChoco(prices, money)
	fmt.Printf("result = %v\n", result)
}

