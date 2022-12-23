package main
import (
	"fmt"
)

func maxProfit(prices []int) int {
    doNothing := int(0)
    buy := int(-100000)
    sell := int(-100000)
    tmp := int(0)
    for i := 0; i < len(prices); i++ {
        tmp = doNothing
        doNothing = max(doNothing, sell)
        buy = max(buy, tmp - prices[i])
        sell = buy + prices[i]
    }

    return max(doNothing, sell)
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 3
	// prices := []int{1,2,3,0,2}

	// result: 0
	prices := []int{1}

	// result: 
	// prices := []int{}

	result := maxProfit(prices)
	fmt.Printf("result = %v\n", result)
}

