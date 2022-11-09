package main
import (
	"fmt"
)

const MAX_NUM_OF_PRICES = 10000

type StockSpanner struct {
	Prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Prices: make([]int, 0, MAX_NUM_OF_PRICES),
	}
}

func (this *StockSpanner) Next(price int) int {
	result := int(1)
	for i := len(this.Prices) - 1; i >= 0; i-- {
		if this.Prices[i] <= price {
			result++
		} else {
			break
		}
	}
	this.Prices = append(this.Prices, price)
	return result
}

func main() {
	obj := Constructor()

	// result: [null, 1, 1, 1, 2, 1, 4, 6]
	prices := []int{100,80,60,70,60,75,85}

	// result: [null, 1, 2, 3, 4, 5]
	// prices := []int{10,10,10,10,50}

	// result: 
	// prices := []int{}

	for _, price := range prices {
		fmt.Printf("obj.Next(%v) = %v\n", price, obj.Next(price))
	}
}

