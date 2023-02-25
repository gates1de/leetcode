package main
import (
	"fmt"
	"sort"
)

type Stock struct {
	Price int
	Index int
}

func maxProfit(prices []int) int {
	n := len(prices)
	ascendings := make([]Stock, n)
	for i, price := range prices {
		ascendings[i] = Stock{Price: price, Index: i}
	}
	sort.Slice(ascendings, func(a, b int) bool {
		return ascendings[a].Price < ascendings[b].Price
	})

	result := int(0)
	i := int(0)
	j := n - 1
	for i < j {
		minStock := ascendings[i]
		maxStock := ascendings[j]
		fmt.Println(minStock, maxStock)
		if minStock.Index < maxStock.Index {
			profit := maxStock.Price - minStock.Price
			if profit > result {
				result = profit
			} else {
				break
			}
		} else {
			j--
			if i >= j {
				i++
				j = n - 1
			}
		}
	}
	return result
}

func main() {
	// result: 5
	// prices := []int{7,1,5,3,6,4}

	// result: 0
	prices := []int{7,6,4,3,1}

	// result: 
	// prices := []int{}

	result := maxProfit(prices)
	fmt.Printf("result = %v\n", result)
}

